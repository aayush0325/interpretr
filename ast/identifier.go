package ast

import "interpretr/token"

// Implements the `Expression` interface
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal } // also returns the variable name
func (i *Identifier) String() string       { return i.Value }         // return the variable name

/*

EXAMPLE USAGE

	&ast.Identifier{
		Token: token.Token{
			Type:    token.IDENT,
			Literal: "myVar",
		},
		Value: "myVar",
	},

*/
