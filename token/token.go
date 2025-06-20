package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"for":    FOR,
	"while":  WHILE,
}

const ILLEGAL = "ILLEGAL"
const EOF = "EOF"

const IDENT = "IDENT"
const INT = "INT"

const ASSIGN = "="
const ADD = "+"
const SUBTRACT = "-"
const DIVIDE = "/"
const MULTIPLY = "*"

const LT = ">"
const LTE = ">="
const GT = "<"
const GTE = "<="
const EQ = "=="
const NEQ = "!="

const NOT = "!"

const COMMA = ","
const SEMICOLON = ";"

const LPAREN = "("
const RPAREN = ")"
const LBRACE = "{"
const RBRACE = "}"

const FUNCTION = "FUNCTION"
const RETURN = "RETURN"

const LET = "LET"

const IF = "IF"
const ELSE = "ELSE"

const TRUE = "TRUE"
const FALSE = "FALSE"

const WHILE = "WHILE"
const FOR = "FOR"
