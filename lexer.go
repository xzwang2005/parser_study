package main

import (
	"strings"
)

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
			return
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

func (it *Lexer) GetNextToken() {
	it.SkipWhitespace()
	// when Done() = true, return nil for currentToken
	if it.Done() {
		return
	}
	currentChar := it.text[it.pos]
	switch currentChar {
	case '+':
		it.Advance()
		it.currentToken = &Token{
			label:   PLUS,
			literal: "+",
		}
	case '-':
		it.Advance()
		it.currentToken = &Token{
			label:   MINUS,
			literal: "-",
		}
	case '*':
		it.Advance()
		it.currentToken = &Token{
			label:   MULT,
			literal: "*",
		}
	case '/':
		it.Advance()
		it.currentToken = &Token{
			label:   DIV,
			literal: "/",
		}
	default:
		if currentChar >= '0' && currentChar <= '9' {
			it.currentToken = it.GetInteger()
		}
	}
}
