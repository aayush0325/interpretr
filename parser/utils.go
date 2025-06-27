package parser

import (
	"interpretr/lexer"
	"interpretr/token"
)

// Returns a parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Reads two tokens to set curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

// Increments both pointers by one
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Compares current token with an input token
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// Compares peek token with an input token
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
