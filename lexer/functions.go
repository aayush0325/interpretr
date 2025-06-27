package lexer

import "interpretr/utils"

type Lexer struct {
	input    string
	position int
	readPos  int
	char     byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0 // EOF
	} else {
		l.char = l.input[l.readPos]
	}

	l.position = l.readPos
	l.readPos += 1
}

// Reads identifiers
func (l *Lexer) readIdent() string {
	i := l.position
	for utils.IsLetter(l.char) {
		l.readChar()
	}
	j := l.position
	return l.input[i:j]
}

// Reads integers in code if encountered
func (l *Lexer) readInt() string {
	i := l.position
	for utils.IsDigit(l.char) {
		l.readChar()
	}
	j := l.position
	return l.input[i:j]
}

// Walk the 2 pointers over any whitespace
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

// Returns the value of the next character
func (l *Lexer) nextChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}
