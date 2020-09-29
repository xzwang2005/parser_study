package main

import (
	"errors"
)

type Interpreter struct {
	Lexer
}

func NewInterpreter(text string) *Interpreter {
	it := Interpreter{
		Lexer{
			text:         text,
			pos:          0,
			currentToken: nil,
		},
	}
	it.GetNextToken()
	return &it
}

func (it *Interpreter) Eat(tokenLabel int) error {
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
	factor:= ('+'|'-')factor | INTEGER | '(' expr ')'

*/

// Factor spit out an integer node or an expr node
func (it *Interpreter) Factor() *AstNode {
	switch it.currentToken.label {
	case LPAREN:
		it.Eat(LPAREN)
		node := it.Expr()
		it.Eat(RPAREN)
		return node
	case INTEGER:
		node := &AstNode{
			token: it.currentToken,
		}
		it.Eat(INTEGER)
		return node
	case PLUS:
		fallthrough
	case MINUS:
		node := &AstNode{
			token: it.currentToken,
		}
		it.Eat(it.currentToken.label)
		leftChild := &AstNode{
			token: &Token{
				label:   INTEGER,
				literal: "0",
			},
		}
		node.left = leftChild
		node.right = it.Factor()
		return node
	default:
		return nil
	}
}

func (it *Interpreter) Term() *AstNode {
	node := it.Factor()
	for {
		if it.currentToken.label == MULT || it.currentToken.label == DIV {
			fNode := &AstNode{
				token: it.currentToken,
			}
			fNode.left = node
			it.Eat(it.currentToken.label)
			fNode.right = it.Factor()
			node = fNode
		} else {
			return node
		}
	}
}

func (it *Interpreter) Expr() *AstNode {

	node := it.Term()

	for {
		if it.currentToken.label == PLUS || it.currentToken.label == MINUS {
			fNode := &AstNode{
				token: it.currentToken,
			}
			fNode.left = node
			it.Eat(it.currentToken.label)
			fNode.right = it.Term()
			node = fNode
		} else {
			return node
		}
	}
}
