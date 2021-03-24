package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinary(t *testing.T) {
	tests := []struct {
		decimal  string
		expected string
	}{
		{decimal: "42", expected: "000000000101010"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, ConvToBinary(test.decimal), test.decimal)
	}
}
