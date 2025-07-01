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

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}
