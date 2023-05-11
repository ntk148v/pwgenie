// Copyright 2023 Kien Nguyen-Tuan <kiennt2609@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func printHelp() {
	var helpText = `
pwgenie is a simple password generator.
<https://github.com/ntk148v/pwgenie>

Usage
-----

  pwgenie [OPTIONS] <SUBCOMMAND> [SUBCOMMAND-OPTIONS]

Options
-------

  -allow-repeat
		Allow repeat characters in the generated password

  -no-clipboard
		Disable automatic copying of generated password to clipboard

Subcommands
-----------

  human    Generate a human-friendly memorable password
  random   Generate a random password with specified complexity
  pin      Generate a random numeric PIN code

Run subcommand with '-h' for subcommand's options.

Example
-------

  $ pwgenie human
  trade clash striking underdog arbitrate

  $ pwgenie human -sep -
  preplan-mousiness-joining-eskimo-linguist

  $ pwgenie random
  bwuelvko

  $ pwgenie random -symb -num -upper
  _U*HkTzA
`
	fmt.Fprintln(os.Stderr, helpText)
	os.Exit(0)
}

func exitOnError(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

var (
	// ErrLettersExceedsAvailable is the error returned with the number of letters
	// exceeds the number of available letters and repeats are not allowed.
	ErrTooManyCharacters = errors.New("number of characters exceeds available letters and repeats are not allowed")
)

func main() {
	allowRepeat := flag.Bool("allow-repeat", false, "Allow repeat characters in the generated password")
	noClipboard := flag.Bool("no-clipboard", false, "Disable automatic copying of generated password to clipboard")
	flag.Usage = printHelp
	flag.Parse()

	// Memorable password
	human := flag.NewFlagSet("human", flag.ExitOnError)
	human.Usage = func() {
		fmt.Fprintf(os.Stderr, "Generate a human-friendly memorable password\n\n")
		fmt.Fprintf(os.Stderr, "Usage of '%s human':\n", os.Args[0])
		human.PrintDefaults()
	}
	words := human.Int("words", 5, "The number of words in the generated password")
	separator := human.String("sep", " ", "The separator for words in the generated password")
	capitalize := human.Bool("cap", false, "Enable capitalization of each word in the generated password")

	// Random
	random := flag.NewFlagSet("random", flag.ExitOnError)
	lenChars := random.Int("length", 8, "The number of characters in the generated password")
	hasUpper := random.Bool("upper", false, "Enable the inclusion of upper-case letters in the generated passwords")
	hasDigits := random.Bool("digit", false, "Enable the inclusion of numbers in the generated password")
	hasSymbols := random.Bool("symbol", false, "Enable the inclusion of symbols in the generated password")
	random.Usage = func() {
		fmt.Fprintf(os.Stderr, "Generate a random password with specified complexity\n\n")
		fmt.Fprintf(os.Stderr, "Usage of '%s random':\n", os.Args[0])
		random.PrintDefaults()
	}

	// Pin
	pin := flag.NewFlagSet("pin", flag.ExitOnError)
	pin.Usage = func() {
		fmt.Fprintf(os.Stderr, "Generate a random numeric PIN code\n\n")
		fmt.Fprintf(os.Stderr, "Usage of '%s pin':\n", os.Args[0])
		pin.PrintDefaults()
	}
	lenNums := pin.Int("length", 6, "The number of digits in the generated PIN code")

	if len(os.Args) < 2 {
		printHelp()
	}

	// Init Rand that uses random values from src
	// to generate other random values.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var (
		pass string
		err  error
	)

	args := flag.Args()
	switch args[0] {
	case "human":
		_ = human.Parse(args[1:])
		pass, err = genHuman(r, *words, *separator, *capitalize, *allowRepeat)
	case "random":
		_ = random.Parse(args[1:])
		pass, err = genRandom(r, *lenChars, *hasUpper, *hasDigits, *hasSymbols, *allowRepeat)
	case "pin":
		_ = pin.Parse(args[1:])
		pass, err = genPIN(r, *lenNums, *allowRepeat)
	default:
		printHelp()
	}

	if err != nil {
		exitOnError(err.Error())
	}

	// Print and copy to clipboard
	if pass != "" {
		fmt.Println(pass)
		if !*noClipboard {
			// Automatically write new pass to clipboard
			_ = clipboard.WriteAll(pass)
		}
	}
}

// genHuman generates a password with the given number of words, separated by the given
// separator.
// If capitalize is true, each word will be capitalized.
func genHuman(r *rand.Rand, words int, separator string, capitalize, allowRepeat bool) (string, error) {
	var (
		formatted []string
		result    string
	)

	if !allowRepeat && words > len(EFFWords) {
		return result, ErrTooManyCharacters
	}

	// Multiple choices from word list
	for i := 0; i < words; i++ {
		word := EFFWords[r.Intn(len(EFFWords))]

		if !allowRepeat && slices.Contains(formatted, word) {
			i--
			continue
		}

		formatted = append(formatted, word)
	}

	// Join the formatted words with the separator
	result = strings.Join(formatted, separator)

	// Capitalize the result if requested
	if capitalize {
		result = cases.Title(language.English).String(result)
	}

	return result, nil
}

// genRandom generates a password with the given number of characters
// using the given character sets.
// This follows Agiles 1Password: https://discussions.agilebits.com/discussion/23842/how-random-are-the-generated-passwords
func genRandom(r *rand.Rand, length int, hasUpper, hasDigits, hasSymbols, allowRepeat bool) (string, error) {
	var (
		maxChars, numLowerChars, numUpperChars, numDigits, numSymbols int
		result                                                        string
	)

	maxChars += len(LowerLetters)
	if hasUpper {
		maxChars += len(UpperLetters)
		numUpperChars = 1
	}

	if hasDigits {
		maxChars += len(Digits)
		numDigits = 1
	}

	if hasSymbols {
		maxChars += len(Symbols)
		numSymbols = 1
	}

	if !allowRepeat && maxChars < length {
		return result, ErrTooManyCharacters
	}

	// calculate characters distributions
	numUpperChars = calcNum(maxChars, len(UpperLetters), length, numUpperChars)
	numDigits = calcNum(maxChars, len(Digits), length, numDigits)
	numSymbols = calcNum(maxChars, len(Symbols), length, numSymbols)

	// The rest is lowercase characters
	numLowerChars = length - numUpperChars - numDigits - numSymbols

	// Lower characters
	for i := 0; i < numLowerChars; i++ {
		ch := randElement(r, LowerLetters)

		if !allowRepeat && strings.Contains(result, ch) {
			i--
			continue
		}

		result = randInsert(r, result, ch)
	}

	// Upper characters
	for i := 0; i < numUpperChars; i++ {
		ch := randElement(r, UpperLetters)

		if !allowRepeat && strings.Contains(result, ch) {
			i--
			continue
		}

		result = randInsert(r, result, ch)
	}

	// Digits
	for i := 0; i < numDigits; i++ {
		ch := randElement(r, Digits)

		if !allowRepeat && strings.Contains(result, ch) {
			i--
			continue
		}

		result = randInsert(r, result, ch)
	}

	// Symbols
	for i := 0; i < numSymbols; i++ {
		ch := randElement(r, Symbols)

		if !allowRepeat && strings.Contains(result, ch) {
			i--
			continue
		}

		result = randInsert(r, result, ch)
	}

	return result, nil
}

// genPIN generates a PIN with the given number of numbers
func genPIN(r *rand.Rand, length int, allowRepeat bool) (string, error) {
	var result string

	if !allowRepeat && length > len(Digits) {
		return result, ErrTooManyCharacters
	}

	// Digits
	for i := 0; i < length; i++ {
		ch := randElement(r, Digits)

		if !allowRepeat && strings.Contains(result, ch) {
			i--
			continue
		}

		result = randInsert(r, result, ch)
	}

	return result, nil
}

// randElement randonly gets an element from given string string
func randElement(r *rand.Rand, s string) string {
	return string(s[r.Intn(len(s))])
}

// randInsert randonly insert an element into given string
func randInsert(r *rand.Rand, s, e string) string {
	if s == "" {
		return e
	}
	pos := r.Intn(len(s) + 1)
	return s[0:pos] + e + s[pos:]
}

// calcNum calculate the number of letters
// based on character distribution in overall.
func calcNum(total, avail, length, initVal int) int {
	var result int

	result = int(math.Floor(float64(avail) / float64(total) * float64(length)))
	if result < 1 {
		result = 1
	}

	return result * initVal
}
