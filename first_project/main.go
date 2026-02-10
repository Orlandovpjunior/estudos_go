package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Println(somar(2, 3))
	a, b := swap(10, 20)
	fmt.Println(a, b)

}

func somar(a, b int) int {
	return a + b
}

// funções privadas começam com letra minusculas, ja as publicas começam com letra maiúscula
// func Foo()
// func foo

// func somar(a int, b int) int {
//	return a+b
//}

func swap(a, b int) (int, int) {
	return b, a
}

// := declarando variável
// = atualizando variável 

func dividir(a, b int) (res int,rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

// retorno pelado, retorna a assinatura da função
func dividir2(a, b int) (res int,rem int) {
	res = a / b
	rem = a % b
	return
}
