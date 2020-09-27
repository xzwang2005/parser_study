package main

import (
	"fmt"
	"strconv"
)

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

var tokenName = []string{"INTEGER", "PLUS", "MINUS", "MULT", "DIV", "LPAREN", "RPAREN", "EOF"}

func (t *Token) String() string {
	return fmt.Sprintf("Token label: %v, literal: %v\n", tokenName[t.label], t.literal)
}

type AstNode struct {
	token *Token
	left  *AstNode
	right *AstNode
}

type EvalFunc func(oprands ...int) int

func IntFunc(ops ...int) int {
	return 0
}

func PlusFunc(ops ...int) int {
	if len(ops) != 2 {
		panic("operands is not 2")
	}
	return ops[0] + ops[1]
}

func MinusFunc(ops ...int) int {
	if len(ops) != 2 {
		panic("operands is not 2")
	}
	return ops[0] - ops[1]
}

func MultFunc(ops ...int) int {
	if len(ops) != 2 {
		panic("operands is not 2")
	}
	return ops[0] * ops[1]
}

func DivFunc(ops ...int) int {
	if len(ops) != 2 {
		panic("operands is not 2")
	}
	if ops[1] == 0 {
		panic("Divided by zero")
	}
	return ops[0] / ops[1]
}

var evalFuncMap = map[int]EvalFunc{
	PLUS:  PlusFunc,
	MINUS: MinusFunc,
	MULT:  MultFunc,
	DIV:   DivFunc,
}

func getEvalFunc() {}

func (n *AstNode) Eval() int {
	if n.token == nil {
		panic("AST node has nil token")
	}
	if n.token.label == INTEGER {
		val, err := strconv.Atoi(n.token.literal)
		if err != nil {
			panic(err)
		}
		return val
	}
	fcn, ok := evalFuncMap[n.token.label]
	if !ok {
		panic("Cannot find evaluation function for token")
	}
	leftVal := n.left.Eval()
	rightVal := n.right.Eval()
	return fcn(leftVal, rightVal)
}
