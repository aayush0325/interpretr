package ast

import "interpretr/token"

// Node to handle statements such as "x+10" in the repl

// Implements the `Statement` interface
type ExpressionStatement struct {
	Token      token.Token // Used for the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
