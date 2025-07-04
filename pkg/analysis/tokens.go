package analysis

type TokenType string

const (
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
	LEFT_ANGLE  TokenType = "["
	RIGHT_ANGLE TokenType = "]"
	QUOTE       TokenType = "\""
	COLON       TokenType = ":"
	COMMA       TokenType = ","
	NULL        TokenType = "null"
	STRING      TokenType = "string"
	NUMBER      TokenType = "number"
	EOF         TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
	Start int
	End   int
	Line  int
}

// type String struct {
// 	Value string
// }
