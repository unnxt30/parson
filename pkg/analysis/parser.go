package analysis

import (
	"fmt"
	"os"
)

type Parser struct {
	Tokens  []Token
	Current int
}

func (p *Parser) Parse() {
	p.declaration()
}

func (p *Parser) declaration() {
	if p.match(LEFT_BRACE) {
		p.content()
		p.consume(RIGHT_BRACE, "Missing '}'")
	}
}

func (p *Parser) content() bool {
	return true
}

func (p *Parser) match(typ TokenType) bool {
	if p.atEnd() {
		return false
	}

	if typ == p.Tokens[p.Current].Type {
		p.Current++
		return true
	}
	return false
}

func (p *Parser) consume(typ TokenType, msg string) {
	if p.atEnd() {
		fmt.Printf("%v\n", msg)
		os.Exit(1)
	}

	if p.peek() != typ {
		fmt.Printf("%v", msg)
		os.Exit(1)
	}

	if !p.atEnd() {
		p.Current++
	}

}

func (p *Parser) peek() TokenType {
	if p.atEnd() {
		os.Exit(1)
	}
	return p.Tokens[p.Current].Type
}

func (p *Parser) atEnd() bool {
	return p.Current >= len(p.Tokens)
}
