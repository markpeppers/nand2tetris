package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/markpeppers/assembler/code"
	"github.com/markpeppers/assembler/parser"
	"github.com/markpeppers/assembler/symbol"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Too few arguments\n")
		os.Exit(1)
	}
	filename := os.Args[1]
	base := strings.Split(filepath.Base(filename), ".")[0]
	rawFilename := fmt.Sprintf("%sL.asm", base)

	// Scan for labels (line numbers)
	p := newParser(filename)
	lineNum := 0
	for p.Advance() {
		label, isLabel := p.Label()
		if isLabel {
			symbol.Table[label] = lineNum
			continue
		}
		lineNum++
	}

	// Replace labels, variables
	outFile := newOutfile(rawFilename)
	p = newParser(filename)
	curValue := 16
	for p.Advance() {
		// Skip labels, eg (HERE)
		if _, isLabel := p.Label(); isLabel {
			continue
		}
		variable, isVariable := p.AInstruction()
		// If instruction, just print it
		if !isVariable {
			fmt.Fprintf(outFile, "%s\n", p.ReturnLine())
			continue
		}
		if !unicode.IsLetter(rune(variable[0])) {
			// Just write it, eg @0
			fmt.Fprintf(outFile, "%s\n", p.ReturnLine())
			continue
		}
		value, present := symbol.Table[variable]
		if !present {
			symbol.Table[variable] = curValue
			value = curValue
			curValue++
		}
		fmt.Fprintf(outFile, "@%d\n", value)
	}
	outFile.Close()

	// Write the final hack file
	hackFilename := fmt.Sprintf("%s.hack", base)
	hackFile := newOutfile(hackFilename)
	parser := newParser(rawFilename)
	for parser.Advance() {
		// fmt.Printf("/%s/\n", parser.ReturnLine())
		if parser.ReturnLine()[0] == '@' {
			fmt.Fprintln(hackFile, parser.Binary())
			continue
		}
		fmt.Fprintf(hackFile, "111%s%s%s\n", code.Comp(parser.Comp()), code.Dest(parser.Dest()), code.Jump(parser.Jump()))
	}
}

func newParser(filename string) *parser.Parser {
	parser, err := parser.NewParser(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating parser: %v\n", err)
		os.Exit(1)
	}
	return parser
}

func newOutfile(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating intermediate file: %v", err)
		os.Exit(1)
	}
	return f
}
