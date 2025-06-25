package parser

import (
	"interpretr/lexer"
	"interpretr/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}
