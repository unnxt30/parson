package main

// here the scanning of files/data and conversion to Tokens takes place.

// a list of tokens converted from the source file/data.
type Scanner struct {
	Tokens []Token
}

func (s *Scanner) Scan(src []byte) {
}

func (s *Scanner) AddToken(token Token) {

}
