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

func (l *Lexer) readIdent() string {
	i := l.position
	for utils.IsLetter(l.char) {
		l.readChar()
	}
	j := l.position
	return l.input[i:j]
}

func (l *Lexer) readInt() string {
	i := l.position
	for utils.IsDigit(l.char) {
		l.readChar()
	}
	j := l.position
	return l.input[i:j]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) nextChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}
