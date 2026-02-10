package firstproject

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr2 := [10]int{5: 400, 7: 300} // [0 0 0 0 0 400 0 300 0 0]

	const x = 10
	arr3 := [x]string{4: "Hello, World"} // [    Hello, World     ]

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
