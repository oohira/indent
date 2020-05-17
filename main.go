package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func indent(s string, count int) string {
	return strings.Repeat(" ", count) + s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(indent(scanner.Text(), 4))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
