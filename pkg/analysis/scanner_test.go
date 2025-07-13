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
		expectError    bool
	}{
		{
			name:           "empty object with empty string",
			source:         `{""}`,
			expectedTokens: 6,
			expectError:    false,
		},
		{
			name:           "simple string",
			source:         `"foo"`,
			expectedTokens: 4,
			expectError:    false,
		},
		{
			name:           "alpha-numeric string",
			source:         `"1foo"`,
			expectedTokens: 4,
			expectError:    false,
		},
		{
			name:           "empty object",
			source:         `{}`,
			expectedTokens: 3,
			expectError:    false,
		},
		{
			name:           "Number",
			source:         `-106`,
			expectedTokens: 2,
			expectError:    false,
		},
		{
			name:           "Number with mathematical constant",
			source:         `-1e2`,
			expectedTokens: 2,
			expectError:    false,
		},
		{
			name:           "number with mathematical constant followed by -",
			source:         `-1e-2`,
			expectedTokens: 2,
			expectError:    false,
		},
		{
			name:           "Positive number",
			source:         `16e5`,
			expectedTokens: 2,
			expectError:    false,
		},
		{
			name:           "valid Decimal Number",
			source:         `1.55`,
			expectedTokens: 2,
			expectError:    false,
		},
		// {
		// 	name:        "invalid decimal Number",
		// 	source:      `1.4.4`,
		// 	expectError: true,
		// },
		{
			name:           "escape character in string",
			source:         `"hello there \"genious\" !"`,
			expectedTokens: 4,
		},
		{
			name:           "valid unicode",
			source:         `"\uFFFF"`,
			expectedTokens: 4,
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
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedTokens, len(tokens))
		})
	}

}
