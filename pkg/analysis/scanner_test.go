package analysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokens(t *testing.T) {
	src :=
		`
		{""}
		`
	scanner := Scanner{
		Tokens:  []Token{},
		Line:    0,
		Current: 0,
		Source:  []byte(src),
	}

	// t.Log(len(src))

	tokens := scanner.Scan()

	assert.Equal(t, len(tokens), 4)

}
