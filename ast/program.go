package ast

import "bytes"

// Implements the `Node` interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Given a prograam, generate a standardized string
func (p *Program) String() string {
	var out bytes.Buffer

	// For each statement in the statements array, call it's `String()` method and write it into the buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
