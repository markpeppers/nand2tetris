package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	scanner *bufio.Scanner
	line    string
}

func NewParser(filename string) (*Parser, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(file)
	return &Parser{scanner: s}, nil
}

func (p *Parser) Advance() bool {
	for p.scanner.Scan() {
		p.line = p.scanner.Text()
		// Remove comments
		commentIndex := strings.Index(p.line, "//")
		if commentIndex > -1 {
			p.line = p.line[:commentIndex]
		}
		// Skip blank lines and remove whitespace
		p.line = strings.Replace(p.line, " ", "", -1)
		if len(p.line) > 0 {
			return true
		}
	}
	return false
}

func (p Parser) Label() (string, bool) {
	if p.line[0] != '(' {
		return "", false
	}
	label := strings.Trim(p.line, "(")
	label = strings.Trim(label, ")")
	return label, true
}

func (p Parser) AInstruction() (string, bool) {
	if p.line[0] != '@' {
		return "", false
	}
	return p.line[1:], true
}

func (p Parser) Comp() string {
	if p.line[0] == '@' {
		return ""
	}
	eq := strings.Index(p.line, "=")
	if eq < 0 {
		eq = 0
	} else {
		eq++
	}
	end := strings.Index(p.line, ";")
	if end < 0 {
		end = len(p.line)
	}
	return p.line[eq:end]
}

func (p Parser) Dest() string {
	if p.line[0] == '@' {
		return ""
	}
	eq := strings.Index(p.line, "=")
	if eq < 0 {
		return ""
	}
	return p.line[:eq]
}

func (p Parser) Jump() string {
	if p.line[0] == '@' {
		return ""
	}
	sem := strings.Index(p.line, ";")
	if sem < 0 {
		return ""
	}
	return p.line[sem+1:]
}

func (p Parser) ReturnLine() string {
	return p.line
}

func (p Parser) Binary() string {
	s := p.line[1:]
	return fmt.Sprintf("0%s", ConvToBinary(s))
}

func ConvToBinary(s string) string {
	binary := ""
	in, _ := strconv.Atoi(s)
	for n := 0; n < 15; n++ {
		div := in / 2
		if in-2*div > 0 {
			binary = fmt.Sprintf("1%s", binary)
		} else {
			binary = fmt.Sprintf("0%s", binary)
		}
		in = div
	}
	return binary
}
