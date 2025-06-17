package main

type TokenType int

const (
	LEFT_BRACE TokenType = iota
	RIGHT_BRACE
	COLON
)

type Token struct {
	Type  TokenType
	Value string
	Start int
	End   int
}

func (t Token) String() string {
	switch t.Type {
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COLON:
		return ":"
	default:
		return t.Value
	}
}
