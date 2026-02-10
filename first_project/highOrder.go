package firstproject

import "fmt"

func main() {

	f := somar(2)
	x := f(1)
	fmt.Println(x)
}

func somar(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
