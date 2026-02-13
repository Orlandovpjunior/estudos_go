package main

import "fmt"

type Orlando struct{}

func (Orlando) String() string {
	return "Esse é um teste"
}

func main() {
	o := Orlando{}
	fmt.Println(o)
}
