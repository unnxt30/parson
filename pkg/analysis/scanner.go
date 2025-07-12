package analysis

import (
	"fmt"
	"os"
	"unicode"
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

func (s *Scanner) Scan() ([]Token, error) {
	if len(s.Source) == 0 {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
	var err error
	for !s.isAtEnd() {
		err = s.scanToken()
	}

	s.addToken(Token{
		Type:  EOF,
		Value: "",
		Start: s.Current,
		End:   s.Current,
		Line:  s.Line,
	})

	return s.Tokens, err
}

func (s *Scanner) scanToken() error {
	var err error
	if s.isAtEnd() {
		return nil
	}
	var token = string(s.Source[s.Current])
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

		err = s.handleString()

	case ":":
		if s.isAtEnd() {
			break
		}

		currToken = Token{
			Type:  COLON,
			Value: ":",
			Start: s.Current,
			End:   s.Current + len(token),
			Line:  s.Line,
		}
		s.addToken(currToken)
		s.Current += len(token)
		s.handleValue()

	case "[":
		if s.isAtEnd() {
			return fmt.Errorf("unexpected end of input at line: %v, col: %v", s.Line, s.Current)
		}
		s.handleArray()

	case "null":
		if s.isAtEnd() {
			return fmt.Errorf("unexpected end of input at line: %v, col: %v", s.Line, s.Current)
		}

	case "true":
		if s.isAtEnd() {
			return fmt.Errorf("unexpected end of input at line: %v, col: %v", s.Line, s.Current)
		}
	case "false":
		if s.isAtEnd() {
			return fmt.Errorf("unexpected end of input at line: %v, col: %v", s.Line, s.Current)
		}
	case "-":
		if s.isAtEnd() {
			return fmt.Errorf("unexpected end of input at line: %v, col: %v", s.Line, s.Current)
		}
		s.Current++
		s.handleNumber(true)
	case "\n":
		s.Line++
		s.Current++
	case " ", "\t", "\r":
		s.Current++
	default:
		if unicode.IsDigit(rune(token[0])) {
			s.handleNumber(false)
		}

	}

	return err

}

func (s *Scanner) addToken(token Token) {
	s.Tokens = append(s.Tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) handleString() error {
	var val string
	start := s.Current
	var endQt *Token
	for !s.isAtEnd() {
		if s.Source[s.Current] == '"' {
			endQt = &Token{
				Type:  QUOTE,
				Value: "\"",
				Start: s.Current,
				End:   s.Current,
				Line:  s.Line,
			}
			s.Current++
			break
		}
		val += string(s.Source[s.Current])
		s.Current++
	}

	end := s.Current - 1

	if s.isAtEnd() && endQt == nil {
		return fmt.Errorf("unterminated string at line: %v", s.Line)
	}

	s.addToken(Token{
		Type:  STRING,
		Value: val,
		Start: start,
		End:   end,
		Line:  s.Line,
	})

	s.addToken(*endQt)

	return nil
}

// handle cases like -12, 1.6e5, -3.14 etc.
// Decimals can be handled recursively
func (s *Scanner) handleNumber(isNegative bool) {
	var num string
	for !s.isAtEnd() && unicode.IsDigit(rune(s.Source[s.Current])) {
		num += string(s.Source[s.Current])
		s.Current++
	}

	if isNegative {
		num = "-" + num
	}

	tok := Token{
		Type:  NUMBER,
		Value: num,
		Start: s.Current - len(num),
		End:   s.Current,
		Line:  s.Line,
	}

	s.addToken(tok)
}

func (s *Scanner) handleArray() {

}

// Keys can only be strings, but values can be strings, number, object, array, true, false, or null.
func (s *Scanner) handleValue() {

}