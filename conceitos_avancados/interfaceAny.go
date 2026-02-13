package main

import (
	"fmt"
	"go/types"

	"golang.org/x/text/cases"
)

type Animal interface {
	Sound() string
}

type Dog struct {
}

type Cat struct {
}

func (Dog) Sound() string {
	return "Mial!"
}

func (Cat) Sound() string {
	return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func (Dog) Interface() {
	fmt.Println("dog interface called")
}

func takeAnimal(a Animal) {
	switch t:= a.(type){
		case *Dog:
			t.Sound()
		case *Cat:
			t.Sound()
	}
	
}

func main() {
	var n any
	n = 10
	n = ""
	n = []int{123}
	fmt.Println(n)
	takeEmptyInterface("")
	takeEmptyInterface(10)
	takeEmptyInterface(Dog{})
	var a Animal
	takeEmptyInterface(a)
}
