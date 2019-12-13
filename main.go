package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	// Lets generate a static family tree as
	// provided in the problem statement
	staticTreeGenerator()
}

func main() {
	// The first argument from command-line
	// Is expected to be the absolute path to the
	// input file
	file, err := os.Open(os.Args[1])
	// Let's close the file after the current execution context
	// in a defered manner
	defer file.Close()
	// We are good to go if there are no errors
	if err == nil {
		fileContent, err := ioutil.ReadAll(file)
		if err == nil {
			lines := strings.Split(string(fileContent), "\n")
			processInputLines(lines)
		}
	}
}

func processInputLines(lines []string) {

}
