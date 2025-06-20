package lexer

import (
	"interpretr/token"
	"interpretr/utils"
)

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.nextChar() == '=' {
			tok.Literal = "=="
			tok.Type = token.EQ
			l.readChar() // to keep the pointers in sync
		} else {
			tok = token.NewToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = token.NewToken(token.ADD, l.char)
	case '-':
		tok = token.NewToken(token.SUBTRACT, l.char)
	case '/':
		tok = token.NewToken(token.DIVIDE, l.char)
	case '*':
		tok = token.NewToken(token.MULTIPLY, l.char)
	case '>':
		tok = token.NewToken(token.GT, l.char)
	case '<':
		tok = token.NewToken(token.LT, l.char)
	case '!':
		if l.nextChar() == '=' {
			tok.Literal = "!="
			tok.Type = token.NEQ
			l.readChar() // to keep the pointers in sync
		} else {
			tok = token.NewToken(token.NOT, l.char)
		}
	case ',':
		tok = token.NewToken(token.COMMA, l.char)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.char)
	case '(':
		tok = token.NewToken(token.LPAREN, l.char)
	case ')':
		tok = token.NewToken(token.RPAREN, l.char)
	case '{':
		tok = token.NewToken(token.LBRACE, l.char)
	case '}':
		tok = token.NewToken(token.RBRACE, l.char)
	case 0: // returned by readChar() on EOF
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if utils.IsLetter(l.char) {
			tok.Literal = l.readIdent()
			tok.Type = token.GetIdentToken(tok.Literal)
			return tok // early return to stop the pointer increment at l.readChar()
		} else if utils.IsDigit(l.char) {
			tok.Literal = l.readInt()
			tok.Type = token.INT
			return tok // early return to stop the pointer increment at l.readChar()
		} else {
			tok = token.NewToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}
