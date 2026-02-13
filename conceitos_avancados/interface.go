package main

import (
	"conceitos_avancados/bar"
	"fmt"
)

type Animal interface {
	Sound() string
}

type Dog struct {
}

func (Dog) Sound() string {
	return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func (Dog) Interface() {
	fmt.Println("dog interface called")
}

func main() {
	dog := Dog{}
	whatDoesThisAnimalSay(dog)
	bar.TakeFoo(dog)
}
