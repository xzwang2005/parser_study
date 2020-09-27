package main

import "fmt"

const (
	// token type
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	MULT    = "MULT"
	DIV     = "DIV"
	EOF     = "EOF"
)

type Token struct {
	label   string
	literal string
}

func (t *Token) String() string {
	return fmt.Sprintf("Token label: %v, literal: %v\n", t.label, t.literal)
}
