package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func sayIt(word string, repetitions int, shout bool) {
	var i int = 0
	if shout {
		word = strings.ToUpper(word)
	}
	for i=0; i<repetitions; i++ {
		fmt.Printf("%s\n", word)
	}
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
