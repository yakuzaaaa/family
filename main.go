package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/family/models"
	"github.com/family/controller"
)

func init() {
	// Lets generate a static family tree as
	// provided in the problem statement
	// This generates the King Shan Family tree
	// And stores it in few in-memory data structures
	models.StaticTreeGenerator()
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
			// Converting bytes to string
			// Spliting the string by line break and getting each query
			lines := strings.Split(string(fileContent), "\n")
			// Process the Array of queries
			processInputLines(lines)
		}
	}
}

func processInputLines(lines []string) {
	for _, line := range lines {
		if len(line) > 0 {
			// Empty line cannot be a valid query
			words := strings.Split(line, " ")
			query := words[0]
			input := words[1:]

			controller.Handle(query, input)
		}
	}
}
