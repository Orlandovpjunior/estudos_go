package main

import (
	"fmt"
)

func main() {

	doDefer()

}

func doDefer() {
	x := 10

	defer func(y int) {
		fmt.Println(x)
	}(x)

	x = 50
	fmt.Println(x)
}

// é uma stack
func doDefer2() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}
