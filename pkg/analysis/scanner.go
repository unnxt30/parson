package analysis

import (
	"fmt"
	"os"
)

// here the scanning of files/data and conversion to Tokens takes place.

// a list of tokens converted from the source file/data.
type Scanner struct {
	Tokens  []Token
	Line    int
	Size    int
	Current int
	Source  []byte
}

func (s *Scanner) Scan() []Token {
	if len(s.Source) == 0 {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
	src := s.Source
	for _, b := range src {
		s.scanToken(string(b))
	}

	return s.Tokens
}

func (s *Scanner) AddToken(token Token, start, end int) {
	s.Tokens = append(s.Tokens, token)
}

func (s *Scanner) scanToken(token string) {
	var currToken Token
	switch token {
	case "{":
		currToken = Token{
			Start: s.Current,
			End:   s.Current + len(token),
			Type:  LEFT_BRACE,
			Value: token,
		}
		s.addToken(currToken)
		s.Current += len(token)
	case "}":
		currToken = Token{
			Start: s.Current,
			End:   s.Current + len(token),
			Type:  RIGHT_BRACE,
			Value: token,
		}
		s.addToken(currToken)
		s.Current += len(token)
	case "\n":
		s.Line++
	case " ":
		break
	default:
	}

}

func (s *Scanner) addToken(token Token) {
	s.Tokens = append(s.Tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}
