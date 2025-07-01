package tests

import (
	"interpretr/ast"
	"interpretr/token"
	"testing"
)

func TestLetAST(t *testing.T) {

	// Good demonstration of how the AST is constructed
	program := &ast.Program{
		Statements: []ast.Statement{
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
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
