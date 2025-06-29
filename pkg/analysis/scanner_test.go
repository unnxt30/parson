package analysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokens(t *testing.T) {
	testCases := []struct {
		name           string
		source         string
		expectedTokens int
	}{
		{
			name:           "empty object with empty string",
			source:         `{""}`,
			expectedTokens: 5,
		},
		{
			name:           "simple string",
			source:         `"foo"`,
			expectedTokens: 4,
		},
		{
			name:           "alpha-numeric string",
			source:         `"1foo"`,
			expectedTokens: 4,
		},
		{
			name:           "empty object",
			source:         `{}`,
			expectedTokens: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scanner := Scanner{
				Tokens:  []Token{},
				Line:    0,
				Current: 0,
				Source:  []byte(tc.source),
			}
			tokens := scanner.Scan()
			assert.Equal(t, tc.expectedTokens, len(tokens))
		})
	}

}
