package main

func Calculate(eq string) int {
	it := Interpreter{
		Lexer{
			text:         eq,
			pos:          0,
			currentToken: nil,
		},
	}
	return it.Expr()
}

func main() {
}
