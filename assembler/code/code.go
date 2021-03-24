package code

import (
	"strings"
)

var jmap = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

// Will replace "A" or "M" with "R", and set a=0 for A, a=1 for M
var cmap = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"R":   "110000",
	"!D":  "001101",
	"!R":  "110001",
	"-D":  "001111",
	"-R":  "110011",
	"D+1": "011111",
	"R+1": "110111",
	"D-1": "001110",
	"R-1": "110010",
	"D+R": "000010",
	"D-R": "010011",
	"R-D": "000111",
	"D&R": "000000",
	"D|R": "010101",
}

func Dest(d string) string {
	code := "000"
	if strings.Index(d, "M") > -1 {
		code = replaceAtIndex(code, '1', 2)
	}
	if strings.Index(d, "D") > -1 {
		code = replaceAtIndex(code, '1', 1)
	}
	if strings.Index(d, "A") > -1 {
		code = replaceAtIndex(code, '1', 0)
	}
	return code
}

func replaceAtIndex(str string, replacement byte, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func Jump(j string) string {
	return jmap[j]
}

func Comp(c string) string {
	a := "0"
	aloc := strings.Index(c, "A")
	if aloc > -1 {
		c = replaceAtIndex(c, 'R', aloc)
	} else {
		mloc := strings.Index(c, "M")
		if mloc > -1 {
			c = replaceAtIndex(c, 'R', mloc)
			a = "1"
		}
	}
	return a + cmap[c]
}
