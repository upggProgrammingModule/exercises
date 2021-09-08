package
main

// There are two bugs in this program.  Try to find them.
// Good luck!

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

// adjustVolume returns a string that is the same
// text as the input string, but with all letters
// capitalized if shout is true.
func adjustVolume(word string, shout bool) string {
	var adjustedText string
	adjustedText = word
	if shout  {
		adjustedText = strings.ToUpper(word)
	}
	return adjustedText
}

// repeatWord will return a string that is equal to the text of word
// repeated "repetitions" times, with each word separated by a space
func repeatWord(word string, repetitions int) string {
	var response string
	var i int

	if repetitions < 0 {
		log.Fatal("Can not repeat a word a negative number of times\n")
	} else if repetitions == 0 {
		return ""
	}

	// we are now dealing with the case that we must
	// repeat the word one, or more, times
	response = word
	for i = 1; i < repetitions; i++ {
		response = fmt.Sprintf("%s %s", response, word)
	}

	return response
}

func sayIt(word string, repetitions int, shout bool) {
	var textToSay string
	word = adjustVolume(word, shout) // word is now uppercase if shout was true
	textToSay = repeatWord(word, repetitions)
	fmt.Printf("%s\n", textToSay)
}

func usage() {
	fmt.Print(
		"parrot\n" +
			"Usage:\n" +
			"parrot anyWord\n" +
			"options:\n")
	flag.PrintDefaults()
}

func main() {
	var expectedNumArgs int = 1
	var repetitions *int = flag.Int("repetitions", 1, "The number of times you would like me to repeat the word.")
	var shoutAtMe *bool = flag.Bool("shoutAtMe", false, "Specifies if I should say the word loudly.")

	flag.Usage = usage
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	if len(flag.Args()) != expectedNumArgs {
		flag.Usage()
		log.Fatalf("Error: expecting %d arguments, but got %d\n", expectedNumArgs, len(flag.Args()))
	}
	inputWord := flag.Arg(0)
	sayIt(inputWord, *repetitions, *shoutAtMe)
}
