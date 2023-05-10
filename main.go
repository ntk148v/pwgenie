package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func printHelp() {
	var helpText = `
pwgenie is a simple password generator.
<https://github.com/ntk148v/pwgenie>

Usage
-----

  pwgenie <SUBCOMMAND> [OPTIONS]

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

func main() {
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
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var pass string

	switch os.Args[1] {
	case "human":
		human.Parse(os.Args[2:])
		pass = genHuman(r, *words, *separator, *capitalize)
	case "random":
		random.Parse(os.Args[2:])
		pass = genRandom(r, *lenChars, *hasUpper, *hasDigits, *hasSymbols)
	case "pin":
		pin.Parse(os.Args[2:])
		pass = genPIN(r, *lenNums)
	default:
		printHelp()
	}

	// Print and copy to clipboard
	if pass != "" {
		fmt.Println(pass)
		// Automatically write new pass to clipboard
		clipboard.WriteAll(pass)
	}
}

// genHuman generates a password with the given number of words, separated by the given
// separator.
// If capitalize is true, each word will be capitalized.
func genHuman(r *rand.Rand, words int, separator string, capitalize bool) string {
	var (
		formatted []string
		result    string
	)
	// Multiple choices from word list
	for i := 0; i < words; i++ {
		formatted = append(formatted, EFFWords[r.Intn(len(EFFWords))])
	}

	// Join the formatted words with the separator
	result = strings.Join(formatted, separator)

	// Capitalize the result if requested
	if capitalize {
		result = cases.Title(language.English).String(result)
	}

	return result
}

// genRandom generates a password with the given number of characters
// using the given character sets.
// This follows Agiles 1Password: https://discussions.agilebits.com/discussion/23842/how-random-are-the-generated-passwords
func genRandom(r *rand.Rand, length int, hasUpper, hasDigits, hasSymbols bool) string {
	var (
		numLowerChars, numUpperChars, numDigits, numSymbols int
		result                                              string
	)

	if hasUpper {
		// Randomly choice number of upper characters in result
		// At least one upper character should be included in result
		numUpperChars = r.Intn(length-4) + 1
	}

	if hasDigits {
		// Randomly choice number of digits in result
		// At least one digit should be included in result
		numDigits = r.Intn(length-numUpperChars-3) + 1
	}

	if hasSymbols {
		// Randomly choice number of symbols in result
		// At least one symbol should be included in result
		numSymbols = r.Intn(length-numDigits-numUpperChars-2) + 1
	}

	// the rest are lower character
	numLowerChars = length - numDigits - numSymbols - numUpperChars

	// Lower characters
	for i := 0; i < numLowerChars; i++ {
		ch := randElement(r, LowerLetters)
		// TODO(kiennt2609): Check repeat!
		result = randInsert(r, result, ch)
	}

	// Upper characters
	for i := 0; i < numUpperChars; i++ {
		ch := randElement(r, UpperLetters)
		// TODO(kiennt2609): Check repeat!
		result = randInsert(r, result, ch)
	}

	// Digits
	for i := 0; i < numDigits; i++ {
		ch := randElement(r, Digits)
		// TODO(kiennt2609): Check repeat!
		result = randInsert(r, result, ch)
	}

	// Symbols
	for i := 0; i < numSymbols; i++ {
		ch := randElement(r, Symbols)
		// TODO(kiennt2609): Check repeat!
		result = randInsert(r, result, ch)
	}

	return result
}

// genPIN generates a PIN with the given number of numbers
func genPIN(r *rand.Rand, num int) string {
	var result string

	// Digits
	for i := 0; i < num; i++ {
		ch := randElement(r, Digits)
		// TODO(kiennt2609): Check repeat!
		result = randInsert(r, result, ch)
	}

	return result
}

// randElement randonly gets an element from given string string
func randElement(r *rand.Rand, s string) string {
	return string(s[r.Intn(len(s))])
}

// randInsert randonly insert an element into given string
func randInsert(r *rand.Rand, s, e string) string {
	pos := r.Intn(len(s) + 1)
	return s[0:pos] + e + s[pos:]
}
