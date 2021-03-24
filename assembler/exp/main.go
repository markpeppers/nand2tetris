package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Too few arguments\n")
		os.Exit(1)
	}
	filename := os.Args[1]
	bytes, _ := ioutil.ReadFile(filename)
	reader := strings.NewReader(string(bytes))
	scanner := bufio.NewScanner(reader)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] != '(' {
			lineCount++
			continue
		}
		fmt.Println(line, lineCount)
	}

	scanner = bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
