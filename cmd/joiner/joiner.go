package main

import (
	"flag"
	"fmt"
	"github.com/vertgenlab/gonomics/fileio"
	"log"
	"strings"
)

func splitAwayFirstCol(line string, columnDelimiter string) (string, string) {
	var columns []string
	var first string
	var everythingElse string
	columns = strings.Split(line, columnDelimiter)
	first = columns[0]
	everythingElse = strings.Join(columns[1:], columnDelimiter)
	return first, everythingElse
}

func processTwo(inputTwoFilename string, columnDelimiter string) map[string]string {
	var inFile *fileio.EasyReader = fileio.EasyOpen(inputTwoFilename)
	var line string
	var doneReading bool

	//start processTwo
	var first string
	var everythingElse string
	var processTwoMap map[string]string
	processTwoMap = make(map[string]string,0)
	//end processTwo

	for line, doneReading = fileio.EasyNextRealLine(inFile); !doneReading; line, doneReading = fileio.EasyNextRealLine(inFile) {
		//start processTwo
		first, everythingElse = splitAwayFirstCol(line, columnDelimiter)
		processTwoMap[first] = everythingElse
		//end processTwo
	}

	return processTwoMap
}

func processOne(inputOneFilename string, outputFilename string, columnDelimiter string, processTwoMap map[string]string) {
	var inFile *fileio.EasyReader = fileio.EasyOpen(inputOneFilename)
	var outFile *fileio.EasyWriter = fileio.EasyCreate(outputFilename)
	var err error
	var line string
	var doneReading bool
	
	//start processOne
	var lineProcessed string
	var first string
	var Two_everythingElse string
	var found bool
	//end processOne

	for line, doneReading = fileio.EasyNextRealLine(inFile); !doneReading; line, doneReading = fileio.EasyNextRealLine(inFile) {
		//start processOne
		first, _ = splitAwayFirstCol(line, columnDelimiter)
		Two_everythingElse, found = processTwoMap[first]
		if !found {
			fmt.Printf("The key %s does not exist in the 2nd input file!\n", first)
		}
		lineProcessed = line + columnDelimiter + Two_everythingElse
		fmt.Fprintf(outFile, "%s\n", lineProcessed)
		//end processOne
	}
	
	err = inFile.Close()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	err = outFile.Close()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

func joiner(inputOneFilename string, inputTwoFilename string, outputFilename string, columnDelimiter string) {
	processTwoMap := processTwo(inputTwoFilename, columnDelimiter)
	processOne(inputOneFilename, outputFilename, columnDelimiter, processTwoMap)
}

func usage() {
	fmt.Print(
		"joiner\n" +
			"Usage:\n" +
			"joiner inputOne.txt inputTwo.txt output.txt\n" +
			"notes:\n" +
			"options:\n" + 
			"columnDelimiter\n")
	flag.PrintDefaults()
}

func main() {
	var expectedNumArgs int = 3
	var columnDelimiter *string = flag.String("delimiter", "\t", "The string that delimits columns in the input")

	flag.Usage = usage
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	if len(flag.Args()) != expectedNumArgs {
		flag.Usage()
		log.Fatalf("Error: expecting %d arguments, but got %d\n", expectedNumArgs, len(flag.Args()))
	}
	inputOneFilename := flag.Arg(0)
	inputTwoFilename := flag.Arg(1)
	outputFilename := flag.Arg(2)

	joiner(inputOneFilename, inputTwoFilename, outputFilename, *columnDelimiter)
}
