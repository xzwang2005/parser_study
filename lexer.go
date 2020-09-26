package main

import "strings"

type Lexer struct {
	text         string
	pos          int
	currentToken *Token
}

func (it *Lexer) Advance() {
	if !it.Done() {
		it.pos++
	}
}

func (it *Lexer) Done() bool {
	return it.pos >= len(it.text)
}

func (it *Lexer) SkipWhitespace() {
	for {
		if it.Done() || it.text[it.pos] != ' ' {
			break
		}
		it.Advance()
	}
}

func (it *Lexer) GetInteger() *Token {
	var sb strings.Builder
	for {
		if it.Done() || it.text[it.pos] < '0' || it.text[it.pos] > '9' {
			break
		}
		sb.WriteString(string(it.text[it.pos]))
		it.Advance()
	}

	return &Token{
		label:   INTEGER,
		literal: sb.String(),
	}
}

func (it *Lexer) GetNextToken() *Token {
	for {
		if it.Done() {
			break
		}
		currentChar := it.text[it.pos]

		if currentChar == ' ' {
			it.SkipWhitespace()
		}

		if currentChar >= '0' && currentChar <= '9' {
			return it.GetInteger()
		}

		if currentChar == '+' {
			it.Advance()
			return &Token{
				label:   PLUS,
				literal: "+",
			}
		}

		if currentChar == '-' {
			it.Advance()
			return &Token{
				label:   MINUS,
				literal: "-",
			}
		}
	}

	// when Done() = true, return nil for currentToken
	return nil
}
