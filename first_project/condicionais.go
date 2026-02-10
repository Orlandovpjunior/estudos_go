package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	if err := doError(); err != nil {
		fmt.Println(err)
	}

	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x > 0 {
		fmt.Println("Maior que zero")
	}
}

func doError() error {
	return errors.New("Error")
}
