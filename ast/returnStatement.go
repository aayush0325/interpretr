package ast

import "interpretr/token"

// Implements the `Statement` interface
type ReturnStatement struct {
	Token       token.Token // return keyword
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
