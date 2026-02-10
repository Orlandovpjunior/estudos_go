package main

import "fmt"

func main() {
	var x *int
	take(x)
	fmt.Println(x)
}

func take(x *int) {
	*x = 100
}

func create() *int {
	x := 10
	return &x
}

func foo(x *int) {
	*x = 100
}

// *int = ponteito, *x = derreferencia
// quando precisar alterar o valor da variável original, use ponteiros
// panic: runtime error: invalid memory address or nil pointer dereference = fez a derreferencia em um ponteiro nulo (linhas 6 e 7) o go run da erro
