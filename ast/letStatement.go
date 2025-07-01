package ast

import (
	"bytes"
	"interpretr/token"
)

// Implements the `Statement` interface
type LetStatement struct {
	Token token.Token // let keyword
	Name  *Identifier
	Value Expression
}

// dummy function
func (ls *LetStatement) statementNode() {}

// returns "let"
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Given a let statement generate a string such as "let x = y"
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	// Constructs the string "let <variable_name> = "
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	// add the value of the expression to the string
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	// add final ";"
	out.WriteString(";")
	return out.String()
}

/*

EXAMPLE USAGE

	&ast.LetStatement{
		Token: token.Token{
			Type:    token.LET,
			Literal: "let",
		},
		Name: &ast.Identifier{
			Token: token.Token{
				Type:    token.IDENT,
				Literal: "myVar",
			},
			Value: "myVar",
		},
		Value: &ast.Identifier{
			Token: token.Token{
				Type:    token.IDENT,
				Literal: "anotherVar",
			},
			Value: "anotherVar",
		},
	},
*/
