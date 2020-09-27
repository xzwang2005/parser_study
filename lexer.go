package main

import (
	"strings"
	"unicode"
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
		if it.Done() || !unicode.IsSpace(rune(it.text[it.pos])) {
			return
		}
		it.Advance()
	}
}

func (it *Lexer) GetInteger() *Token {
	var sb strings.Builder
	for {
		if it.Done() || !unicode.IsDigit(rune(it.text[it.pos])) {
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
	case '(':
		it.Advance()
		it.currentToken = &Token{
			label:   LPAREN,
			literal: "(",
		}
	case ')':
		it.Advance()
		it.currentToken = &Token{
			label:   RPAREN,
			literal: ")",
		}
	default:
		if unicode.IsDigit(rune(currentChar)) {
			it.currentToken = it.GetInteger()
		}
	}
}
