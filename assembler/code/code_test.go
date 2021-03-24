package code

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDest(t *testing.T) {
	tests := []struct {
		dests    string
		expected string
	}{
		{dests: "", expected: "000"},
		{dests: "M", expected: "001"},
		{dests: "D", expected: "010"},
		{dests: "MD", expected: "011"},
		{dests: "A", expected: "100"},
		{dests: "AM", expected: "101"},
		{dests: "AD", expected: "110"},
		{dests: "AMD", expected: "111"},
		{dests: "ADM", expected: "111"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Dest(test.dests), test.dests)
	}
}

func TestJump(t *testing.T) {
	tests := []struct {
		j        string
		expected string
	}{
		{j: "", expected: "000"},
		{j: "JGT", expected: "001"},
		{j: "JEQ", expected: "010"},
		{j: "JGE", expected: "011"},
		{j: "JLT", expected: "100"},
		{j: "JNE", expected: "101"},
		{j: "JLE", expected: "110"},
		{j: "JMP", expected: "111"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Jump(test.j), test.j)
	}
	fmt.Println(tests)
}

func TestComp(t *testing.T) {
	tests := []struct {
		c        string
		expected string
	}{
		{c: "0", expected: "0101010"},
		{c: "1", expected: "0111111"},
		{c: "-1", expected: "0111010"},
		{c: "A", expected: "0110000"},
		{c: "M", expected: "1110000"},
		{c: "!D", expected: "0001101"},
		{c: "!A", expected: "0110001"},
		{c: "!M", expected: "1110001"},
		{c: "-D", expected: "0001111"},
		{c: "-A", expected: "0110011"},
		{c: "-M", expected: "1110011"},
		{c: "D+1", expected: "0011111"},
		{c: "A-1", expected: "0110010"},
		{c: "M-1", expected: "1110010"},
		{c: "D+A", expected: "0000010"},
		{c: "D+M", expected: "1000010"},
		{c: "D-A", expected: "0010011"},
		{c: "D-M", expected: "1010011"},
		{c: "A-D", expected: "0000111"},
		{c: "M-D", expected: "1000111"},
		{c: "D&A", expected: "0000000"},
		{c: "D&M", expected: "1000000"},
		{c: "D|A", expected: "0010101"},
		{c: "D|M", expected: "1010101"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Comp(test.c), test.c)
	}
}
