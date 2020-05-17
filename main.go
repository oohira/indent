package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	spaces := flag.Int("n", 4, "number of spaces")
	flag.Parse()

	err := indent(os.Stdin, os.Stdout, *spaces)
	if err != nil {
		log.Fatal(err)
	}
}
