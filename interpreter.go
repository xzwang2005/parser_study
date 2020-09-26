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
