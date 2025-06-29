package analysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	src :=
		`{"}`
	scanner := Scanner{
		Tokens:  []Token{},
		Line:    0,
		Current: 0,
		Source:  []byte(src),
	}

	tokens := scanner.Scan()

	parser := Parser{
		Current: 0,
		Tokens:  tokens,
	}

	err := parser.Parse()

	assert.Nil(t, err)

}
