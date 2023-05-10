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
	characters := random.Int("chars", 8, "The number of characters in the generated password")
	hasUpper := random.Bool("upper", false, "Enable the inclusion of upper-case letters in the generated passwords")
	hasNum := random.Bool("num", false, "Enable the inclusion of numbers in the generated password")
	hasSymb := random.Bool("symb", false, "Enable the inclusion of symbols in the generated password")
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
	numbers := pin.Int("num", 6, "The number of digits in the generated PIN code")

	if len(os.Args) < 2 {
		printHelp()
	}

	var pass string

	switch os.Args[1] {
	case "human":
		human.Parse(os.Args[2:])
		pass = genHuman(*words, *separator, *capitalize)
	case "random":
		random.Parse(os.Args[2:])
		pass = genRandom(*characters, *hasUpper, *hasNum, *hasSymb)
	case "pin":
		pin.Parse(os.Args[2:])
		pass = genPIN(*numbers)
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
func genHuman(words int, separator string, capitalize bool) string {
	var (
		formatted []string
		pass      string
	)
	// Multiple choices from word list
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < words; i++ {
		formatted = append(formatted, EFFWords[r.Intn(len(EFFWords))])
	}

	// Join the formatted words with the separator
	pass = strings.Join(formatted, separator)

	// Capitalize the result if requested
	if capitalize {
		pass = cases.Title(language.English).String(pass)
	}

	return pass
}

// genRandom generates a password with the given number of characters
// using the given character sets.
func genRandom(chars int, hasUpper, hasNum, hasSymb bool) string {
	var letters string
	letters = LowerLetters
	if hasUpper {
		letters += UpperLetters
	}

	if hasSymb {
		letters += SymbolLetters
	}

	if hasNum {
		letters += NumberLetters
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := make([]byte, chars)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

// genPIN generates a PIN with the given number of numbers
func genPIN(num int) string {
	letters := NumberLetters
	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := make([]byte, num)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}
