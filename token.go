package main

import "fmt"

const (
	// token type
	INTEGER = 0
	PLUS    = 1
	MINUS   = 2
	MULT    = 3
	DIV     = 4
	LPAREN  = 5
	RPAREN  = 6
	EOF     = 7
)

type Token struct {
	label   int
	literal string
}

func (t *Token) String() string {
	names := []string{"INTEGER", "PLUS", "MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "EOF"}
	return fmt.Sprintf("Token label: %v, literal: %v\n", names[t.label], t.literal)
}
