package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	// token type
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	EOF     = "EOF"
)

type Token struct {
	label   string
	literal string
}

func (t *Token) String() string {
	return fmt.Sprintf("Token label: %v, literal: %v\n", t.label, t.literal)
}

type Interpreter struct {
	text         string
	pos          int
	currentToken *Token
}

func (it *Interpreter) GetNextToken() *Token {
	if it.pos >= len(it.text) {
		return nil
	}
	currentChar := it.text[it.pos]
	if currentChar >= '0' && currentChar <= '9' {
		it.pos++
		return &Token{
			label:   INTEGER,
			literal: string(currentChar),
		}
	}

	if currentChar == '+' {
		it.pos++
		return &Token{
			label:   PLUS,
			literal: "+",
		}
	}
	return nil
}

func (it *Interpreter) Eat(tokenLabel string) error {
	if it.currentToken.label == tokenLabel {
		fmt.Printf("process token: %v\n", it.currentToken)
		it.currentToken = it.GetNextToken()
		return nil
	}
	return errors.New("Error in consuming token")
}

func (it *Interpreter) Expr() int {
	it.currentToken = it.GetNextToken()
	left, err := strconv.Atoi(it.currentToken.literal)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	it.Eat(INTEGER)

	//_ := it.currentToken
	it.Eat(PLUS)

	right, err := strconv.Atoi(it.currentToken.literal)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	it.Eat(INTEGER)

	return left + right
}

func Calculate(eq string) int {
	it := Interpreter{
		text:         eq,
		pos:          0,
		currentToken: nil,
	}
	return it.Expr()
}

func main() {
}
