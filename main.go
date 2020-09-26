package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// token type
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
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

func (it *Interpreter) Advance() {
	if !it.Done() {
		it.pos++
	}
}

func (it *Interpreter) Done() bool {
	return it.pos >= len(it.text)
}

func (it *Interpreter) SkipWhitespace() {
	for {
		if it.Done() || it.text[it.pos] != ' ' {
			break
		}
		it.Advance()
	}
}

func (it *Interpreter) GetInteger() *Token {
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

func (it *Interpreter) GetNextToken() *Token {
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

func (it *Interpreter) Eat(tokenLabel string) error {
	if it.currentToken.label == tokenLabel {
		//fmt.Printf("process token: %v\n", it.currentToken)
		it.currentToken = it.GetNextToken()
		return nil
	}
	return errors.New("Error in consuming token")
}

func (it *Interpreter) Term() (int, error) {
	val, err := strconv.Atoi(it.currentToken.literal)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	it.Eat(INTEGER)
	return val, nil
}

func (it *Interpreter) Expr() int {
	it.currentToken = it.GetNextToken()

	val, _ := it.Term()

	for {
		if it.currentToken == nil {
			break
		}

		if it.currentToken.label == PLUS {
			it.Eat(PLUS)
			right, _ := it.Term()
			val += right
			continue
		}

		if it.currentToken.label == MINUS {
			it.Eat(MINUS)
			right, _ := it.Term()
			val -= right
			continue
		}
	}
	return val
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
