package main

import (
	"flag"
	"fmt"
	"github.com/vertgenlab/gonomics/fileio"
	"log"
	"strings"
	"strconv"
)

// commaListToInts will take in a string of comma
// separated numbers and return a slice of ints
// representing the numbers in the input string
func commaListToInts(commaList string) []int {
	var charlist []string
	charlist = make([]string, len(commaList))
	var intlist []int
	var i int
	charlist = strings.Split(commaList, ",")
	intlist = make([]int, len(charlist))
	for i = 0; i < len(charlist); i++{
		intlist[i],_  = strconv.Atoi(charlist[i])
	}
	return intlist
}

// oneBasedToZeroBased will take a list of numbers in 1-based
// counting and return the same list of numbers converted to 0-based
// counting by substracting one from each number in the slice
func oneBasedToZeroBased(inputNumbers []int) []int {
	var newlist []int
	newlist = make([]int, len(inputNumbers))
	var i int
	for i = range inputNumbers{
		newlist[i] = inputNumbers[i] - 1
	}
	return newlist
}

// reorderColumns will take in the line of a file, split it into columns
// based on delimiter, and then put a subset of the columns back into
// a string based on the order of their indices in columnOrder
// columnOrder is zero-based with the first field having an index of zero
func reorderColumns(line string, delimiter string, columnOrder []int) string {
	// TODO: write this function

	//split the string based on the delimiter

	//build a new string while looping through columnOrder

	//return the new string you built
	var strlist []string
	strlist = make([]string, len(line))
	strlist = strings.Split(line, delimiter)
	var i int
	var result = strlist[columnOrder[0]]
	for i = 1;i < len(columnOrder);i++{
	result = fmt.Sprintf("%s%s%s",result,delimiter,strlist[columnOrder[i]])
	}
	return result
}

func reorder(inputFilename string, columnDelimiter string, commaListOfFields string, outputFilename string) {
	var inFile *fileio.EasyReader = fileio.EasyOpen(inputFilename)
	var outFile *fileio.EasyWriter = fileio.EasyCreate(outputFilename)
	var err error
	var line string
	var doneReading bool
	var collist []int
	collist = make([]int,len(commaListOfFields)) 
	// TODO: rite logic to turn the string of 1-based comma-separated indices
	// into a slice of 0-based ints
	collist = commaListToInts(commaListOfFields)
	collist = oneBasedToZeroBased(collist) 
	for line, doneReading = fileio.EasyNextRealLine(inFile); !doneReading; line, doneReading = fileio.EasyNextRealLine(inFile) {
		// TODO: write the code for inside this loop that parses each line of the input file
		// you will want to use the reorderColumns function above to modify the line you just read in
		fmt.Fprintf(outFile, "%s\n", reorderColumns(line, columnDelimiter, collist))
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

func usage() {
	fmt.Print(
		"reorder\n" +
			"Usage:\n" +
			"reorder input.txt columnNumbers output.txt\n" +
			"notes:\n" +
			" columnNumbers is a comma separated list of the columns you would like" +
			"  in the output file.  It is similar to the -f option in cut, but allows" +
			"  for the reordering of columns" +
			"options:\n")
	flag.PrintDefaults()
}

func main() {
	var expectedNumArgs int = 3
	var delim *string = flag.String("delimiter", "\t", "The string that delimits columns in the input")

	flag.Usage = usage
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	if len(flag.Args()) != expectedNumArgs {
		flag.Usage()
		log.Fatalf("Error: expecting %d arguments, but got %d\n", expectedNumArgs, len(flag.Args()))
	}
	inputFilename := flag.Arg(0)
	listOfFields := flag.Arg(1)
	outputFilename := flag.Arg(2)

	reorder(inputFilename, *delim, listOfFields, outputFilename)
}
