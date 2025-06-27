package parser

import (
	"interpretr/lexer"
	"interpretr/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

// Returns a parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Reads two tokens to set curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}
