package main

import "fmt"

func main() {
	m := map[string]string{
		"Orlando": "Junior",
		"Joaquim": "Pedro",
	}
	//m["Orlando"] =  "Junior"
	// valor, ok := m["Orlando"] // ok fala se o valor existe ou nao
	for k := range m {

		if k == "Orlando" {
			delete(m, k)
		}

	}
	fmt.Println(m) // map[Joaquim:Pedro]
}
