package main

import (
	"fmt"
	"os"

	"github.com/cahfofpai/birthday.md/internal/ics"
	"github.com/cahfofpai/birthday.md/internal/parser"
)

func main() {
	// Parse command line arguments
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Input file '%s' does not exist\n", inputFile)
		os.Exit(1)
	}

	// Parse the input file
	p := parser.NewParser(inputFile)
	birthdays, err := p.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// Print any parsing errors
	errors := p.GetErrors()
	for _, err := range errors {
		fmt.Fprintf(os.Stderr, "Warning: %s\n", err)
	}

	// Generate the ICS file
	g := ics.NewGenerator(outputFile, birthdays)
	err = g.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// Print success message
	fmt.Printf("Successfully converted %d birthdays to ICS format.\n", len(birthdays))
	fmt.Printf("Output written to: %s\n", outputFile)
}

func printUsage() {
	fmt.Println("Syntax: birthday-md <input file> <output file>")
	fmt.Println()
	fmt.Println("Example: birthday-md birthdays.md birthdays.ics")
}