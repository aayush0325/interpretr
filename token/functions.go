package token

// Initializes a new token
func NewToken(TokenType TokenType, Literal byte) Token {
	return Token{
		Type:    TokenType,
		Literal: string(Literal),
	}
}

// Returns the keyword token if matches OR the identifer token
func GetIdentToken(ident string) TokenType {
	if value, ok := keywords[ident]; ok {
		return value
	}
	return IDENT
}
