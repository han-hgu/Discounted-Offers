package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func usage() string {
	//TODO
	usage := "usage: "
	return usage
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(usage())
	}

	rs, err := parseFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs)
	process(rs, runtime.NumCPU())
}
