package main

import (
	"flag"
	"fmt"
	"github.com/vertgenlab/gonomics/fileio"
	"log"
		"strings"
		
)

// commaListToInts will take in a string of comma
// separated numbers and return a slice of ints
// representing the numbers in the input string
func commaListToInts(commaList string) []int {
	
	commaSlice := strings.Split(commaList, ",")
	var new_list []int
	new_list =make([]int, len(commaSlice))
	var i int
	for i=0; i < len (commaSlice); i++ {
	n,_ := strconv.Atoi (commaSlice [i])
	new_list[i] = n
	}
	return new_list

}


// oneBasedToZeroBased will take a list of numbers in 1-based
// counting and return the same list of numbers converted to 0-based
// counting by substracting one from each number in the slice
func oneBasedToZeroBased(inputNumbers []int) []int {
	// TODO: write this function

	// loop through the list of input numbers and
	// build a new list by subtracting one from each.
	var zero_based []int
	zero_based = make([]int,len(inputNumbers))
	var i int
	for i = 0; i < len(inputNumbers); i++{
	zero_based[i] = inputNumbers[i] - 1
}
        return zero_based
}


// reorderColumns will take in the line of a file, split it into columns
// based on delimiter, and then put a subset of the columns back ino
// a string based on the order of their indices in columnOrder
// columnOrder is zero-based with the first field having an index of zero
func reorderColumns(line string, delimiter string, columnOrder []int) string {
	// TODO: write this function

	//split the string based on the delimiter
	lineSlice := strings.Split(line, delimiter)
	//build a new string while looping through columnOrder

	//return the new string you built
	var i int
	var new_line string  
	for i=0; i < len(columnOrder); i++ {
	new_line = new_line + lineSlice[columnOrder[i]]
}
        //return the new string you built
        return new_line

}


func reorder(inputFilename string, columnDelimiter string, commaListOfFields string, outputFilename string) {
	var inFile *fileio.EasyReader = fileio.EasyOpen(inputFilename)
	var outFile *fileio.EasyWriter = fileio.EasyCreate(outputFilename)
	var err error
	var line string
	var doneReading bool

	// TODO: write logic to turn the string of 1-based comma-separated indices
	// into a slice of 0-based ints

	for line, doneReading = fileio.EasyNextRealLine(inFile); !doneReading; line, doneReading = fileio.EasyNextRealLine(inFile) {
		// TODO: write the code for inside this loop that parses each line of the input file
		// you will want to use the reorderColumns function above to modify the line you just read in
		fmt.Println(outFile, "%s\n", reorderColumns (line, columnDelimiter, oneBasedToZeroBased(commaListToInts(commaListOfFields))))
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
		"joiner\n" +
			"Usage:\n" +
			"joiner inputOne.txt inputTwo.txt output.txt\n" +
			"notes:\n" +
			"options:\n")
	flag.PrintDefaults()
}

func splitAwayFirstCol(line string) (string, string) {
	var columns []string
	var first string
	var everythingElse string
	columns = strings.Split(line, ":")
	first = columns [0]
	everythingElse = strings.Join(columns [1:], ":")
	return first, everythingElse


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

	joiner(inputOneFilename, inputTwoFilename, outputFilename)

	reorder(inputFilename, *delim, listOfFields, outputFilename)
}

