package main

func Calculate(eq string) int {
	it := NewInterpreter(eq)
	astRoot := it.Expr()
	return astRoot.Eval()
}

func main() {
}
