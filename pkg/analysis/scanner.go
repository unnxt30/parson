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

	for !s.isAtEnd() {
		s.scanToken()
	}

	s.addToken(Token{
		Type:  EOF,
		Value: "",
		Start: s.Current,
		End:   s.Current,
		Line:  s.Line,
	})

	return s.Tokens
}

func (s *Scanner) scanToken() {
	if s.isAtEnd() {
		return
	}
	var token = string(string(s.Source[s.Current]))
	var currToken Token
	switch token {
	case "{":
		currToken = Token{
			Start: s.Current,
			End:   s.Current + len(token),
			Type:  LEFT_BRACE,
			Value: token,
			Line:  s.Line,
		}
		s.addToken(currToken)
		s.Current += len(token)
	case "}":
		currToken = Token{
			Start: s.Current,
			End:   s.Current + len(token),
			Type:  RIGHT_BRACE,
			Value: token,
			Line:  s.Line,
		}
		s.addToken(currToken)
		s.Current += len(token)
	case "\"":
		if s.isAtEnd() {
			break
		}

		currToken = Token{
			Type:  QUOTE,
			Value: "\"",
			Start: s.Current,
			End:   s.Current + len(token),
			Line:  s.Line,
		}

		s.addToken(currToken)
		s.Current += len(token)

	case "\n":
		s.Line++
		s.Current++
	case " ", "\t", "\r":
		s.Current++
	default:
		if s.isString(token) {
			currToken = s.handleString()
			s.addToken(currToken)
			s.Current += len(currToken.Value)
		}
	}

}

func (s *Scanner) addToken(token Token) {
	s.Tokens = append(s.Tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) isString(tok string) bool {
	return (tok >= "a" && tok <= "z") ||
		(tok >= "A" && tok <= "Z")

}

func (s *Scanner) handleString() Token {
	var val string
	start := s.Current
	for s.isString(string(s.Source[s.Current])) {
		val += string(s.Source[s.Current])
		s.Current++
	}
	end := s.Current

	return Token{
		Type:  STRING,
		Value: val,
		Start: start,
		End:   end,
		Line:  s.Line,
	}
}
