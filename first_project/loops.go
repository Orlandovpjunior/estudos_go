package main

import "fmt"

func main() {
	var res int
	for i := 0; i < 10; i++ {
		res++
	}
	fmt.Println(res)

	// While true

	/*
		for   {
			res++
		}
	*/

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, elem := range arr {
		fmt.Println(elem)
	}

	for range 10 {
		fmt.Println("Dentro")
	}

	arr2 := [3] int {1,2,3}

	for i, elem:= range arr2{
		fmt.Println(&i, &elem)
	}

}
