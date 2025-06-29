package analysis

import (
	"fmt"
	"os"
)

type Parser struct {
	Tokens  []Token
	Current int
}

func (p *Parser) Parse() error {
	err := p.value()

	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) value() error {
	if p.match(LEFT_BRACE) {
		err := p.value()
		if err != nil {
			return err
		}

		err = p.consume(RIGHT_BRACE, fmt.Sprintf("Missing '}' at line:%v, column: %v", p.Tokens[p.Current].Line, p.Tokens[p.Current].Start))
		if err != nil {
			return err
		}
	} else if p.match(QUOTE) {
		p.string()
		err := p.consume(QUOTE, fmt.Sprintf("Missing closing quote at line: %v, column: %v", p.Tokens[p.Current].Line, p.Tokens[p.Current].Start))

		if err != nil {
			return err
		}

	}

	return nil
}

func (p *Parser) string() bool {
	p.Current += 1
	return true
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

func (p *Parser) consume(typ TokenType, msg string) error {
	if p.atEnd() {
		return fmt.Errorf("%v", msg)
	}

	if p.peek() != typ {
		return fmt.Errorf("%v", msg)
	}

	p.Current++

	return nil
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
