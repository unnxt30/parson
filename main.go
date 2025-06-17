package main

import (
	"fmt"
	"log"
	"os"
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

	fmt.Println(string(data))

}
