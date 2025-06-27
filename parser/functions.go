package parser

import (
	"interpretr/ast"
	"interpretr/token"
)

// Returns an AST
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}              // Initialize AST
	program.Statements = []ast.Statement{} // Initialize empty array of statements
	for p.curToken.Type != token.EOF {     // while we don't encounter EOF, keep parsing statements
		stmt := p.parseStatement() // Switch based function to parse different types of tokens, LET etc...
		if stmt != nil {
			program.Statements = append(program.Statements, stmt) // Add the statement to the array of statements if true
		}
		p.nextToken()
	}
	return program
}

// Returns a "statement" to be stored in the `program`
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// Function to parse let statements
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// Initialize a let statement with token "LET"
	stmt := &ast.LetStatement{Token: p.curToken}

	// Invalid syntax if the next statement is not an identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Define the variable name in the let statement
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
