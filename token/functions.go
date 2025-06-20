package token

func NewToken(TokenType TokenType, Literal byte) Token {
	return Token{
		Type:    TokenType,
		Literal: string(Literal),
	}
}

func GetIdentToken(ident string) TokenType {
	if value, ok := keywords[ident]; ok {
		return value
	}
	return IDENT
}
