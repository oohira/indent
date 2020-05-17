package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func indent(s string, count int) string {
	return strings.Repeat(" ", count) + s
}

func main() {
	spaces := flag.Int("n", 4, "number of spaces")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(indent(scanner.Text(), *spaces))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
