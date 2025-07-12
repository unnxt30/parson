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
			expectedTokens: 6,
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
		{
			name:           "Number",
			source:         `-106`,
			expectedTokens: 2,
		},
		{
			name:           "Number with mathematical constant",
			source:         `-1e2`,
			expectedTokens: 2,
		},
		{
			name:           "number with mathematical constant followed by -",
			source:         `-1e-2`,
			expectedTokens: 2,
		},
		{
			name:           "Positive number",
			source:         `16e5`,
			expectedTokens: 2,
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
			tokens, err := scanner.Scan()
			assert.NoError(t, err)
			t.Log(tokens)
			assert.Equal(t, tc.expectedTokens, len(tokens))
		})
	}

}
