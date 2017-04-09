package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jehaby/kinpa"
)

var (
	filename = flag.String("filename", "", "Kindle clipping file")
	debug    = flag.Bool("debug", false, "Debug mode")
)

func main() {
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Panicf("error opening kindle file: %v", err)
	}

	parser := kinpa.NewParser(*debug)

	hs, bs, err := parser.ParseClippings(f)
	if err != nil {
		log.Panicf("error parsing clipping file: %v", err)
	}

	fmt.Println(hs.Len(), bs.Len())

	kinpa.PrintHighlights(hs, os.Stdout)

	/*
		highlights := hs.storage

		for i := range highlights {
			fmt.Println(&hs.storage[i], hs.storage[i])
		}

	*/
}
