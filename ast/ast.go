package ast

type Node interface {
	TokenLiteral() string
	String() string // Get the whole program back as a string, for testing purposes
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
