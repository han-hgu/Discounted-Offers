package main

import (
	"log"
	"os"
)

func usage() string {
	usage := "usage: Discounted-Offers input_file"
	return usage
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(usage())
	}

	rs, err := ParseFile(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to parse file %v: %v", os.Args[1], err)
	}

	result := processStrings(rs)
	print(result)
}
