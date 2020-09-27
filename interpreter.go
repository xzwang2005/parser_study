package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Interpreter struct {
	Lexer
}

func (it *Interpreter) Eat(tokenLabel string) error {
	if it.currentToken.label == tokenLabel {
		//fmt.Printf("process token: %v\n", it.currentToken)
		it.GetNextToken()
		return nil
	}
	return errors.New("Error in consuming token")
}

/*

grammar:
	expr := term (('+'|'-')term)*
	term := factor (('*'|'/')factor)*
	factor:= INTEGER

*/

func (it *Interpreter) Factor() (int, error) {
	val, err := strconv.Atoi(it.currentToken.literal)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	it.Eat(INTEGER)
	return val, nil
}

func (it *Interpreter) Term() int {
	val, _ := it.Factor()
	for {
		switch it.currentToken.label {
		case MULT:
			it.Eat(MULT)
			right, _ := it.Factor()
			val *= right
		case DIV:
			it.Eat(DIV)
			right, _ := it.Factor()
			val /= right
		default:
			return val
		}
	}
}

func (it *Interpreter) Expr() int {

	val := it.Term()

	for {
		switch it.currentToken.label {
		case PLUS:
			it.Eat(PLUS)
			right := it.Term()
			val += right
		case MINUS:
			it.Eat(MINUS)
			right := it.Term()
			val -= right
		default:
			return val
		}
	}
}
