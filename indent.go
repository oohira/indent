package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func indent(r io.Reader, w io.Writer, count int) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		_, err := fmt.Fprintln(w, indentLine(scanner.Text(), count))
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func indentLine(s string, count int) string {
	return strings.Repeat(" ", count) + s
}
