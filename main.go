package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	spaces := flag.Int("n", 4, "number of spaces")
	flag.Parse()

	var r io.Reader
	if flag.NArg() == 0 {
		r = os.Stdin
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		r = file
	}
	err := indent(r, os.Stdout, *spaces)
	if err != nil {
		log.Fatal(err)
	}
}
