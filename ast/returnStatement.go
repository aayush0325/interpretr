package ast

import (
	"bytes"
	"interpretr/token"
)

// Implements the `Statement` interface
type ReturnStatement struct {
	Token       token.Token // return keyword
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	// Constructs the string "return <return_value>;"
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

/*

EXAMPLE USAGE

	&ast.ReturnStatement{
		Token: token.Token{
			Type:    token.RETURN,
			Literal: "return",
		},
		ReturnValue: &ast.Identifier{
			Token: token.Token{
				Type:    token.IDENT,
				Literal: "x",
			},
			Value: "x",
		},
	},

*/
