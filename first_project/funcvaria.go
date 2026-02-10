package main

import "fmt"

func main(){

	fmt.Println(somar(10,20,30))

}

func somar(nums ...int) int {
	var out int

	for _, n :=range nums{
		out += n
	}

	return  out
}