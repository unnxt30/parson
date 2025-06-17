package main

type TokenType string

const (
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
)

type Token struct {
	Type  TokenType
	Value string
	Start int
	End   int
}
