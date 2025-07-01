package tests

import (
	"interpretr/ast"
	"interpretr/token"
	"testing"
)

func TestReturnAST(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
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
		},
	}

	if program.String() != "return x;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
