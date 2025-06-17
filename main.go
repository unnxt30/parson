package main

import (
	"log"
	"os"

	ana "github.com/unnxt30/parson/pkg/analysis"
)

func main() {
	args := os.Args
	var path string = ""
	if len(args) < 2 {
		log.Fatal("File path not provided")
		os.Exit(1)
	}

	path = args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not parse file")
		os.Exit(1)
	}

	tokens := []ana.Token{}
	scanner := &ana.Scanner{
		Tokens: tokens,
		Line:   0,
		Size:   len(data),
		Source: data,
	}

	tokens = scanner.Scan()

	parser := &ana.Parser{
		Tokens:  tokens,
		Current: 0,
	}

	parser.Parse()
}
