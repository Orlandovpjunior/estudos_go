package main

import "fmt"

var filmesDb = []string{
	"Matrix",
	"Futebol 1",
	"CR7",
	"Messi",
	"Neymar",
	"Tubarão",
	"Massacre",
	"Bruna surfistinha",
	"Senhor dos aneis",
	"Herman",
	"Manel",
}

func main() {
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // ex ids de filmes favoritos

	filmes := make([]string, 0, 10) // melhor performace para o slice
	// toda lógica do db

	for _, id := range resultsFromApi {
		filme := filmesDb[id]
		fmt.Println(len(filmes), cap(filmes))
		filmes = append(filmes, filme)
		fmt.Println(len(filmes), cap(filmes)) // sempre que o slice precisa alocar mais, ele dobra a capacidade
	}

	fmt.Println(filmes)

}

// full slice expression arr:= [5] int {1,2,3,4,5} silce := arr[:2:2]
