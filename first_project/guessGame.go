package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println(
		"Um número aleatório será sorteado. Tente Acertar. O número é um inteiro entre 0 e 100",
	)

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou! O número sorteado é maior que ", chute)
		case chuteInt > x:
			fmt.Println("Você errou! O número sorteado é menor que ", chute)
		case chuteInt == x:
			fmt.Printf(
				"Você acertou! O número era %d.\n"+
					"Você acertou na tentativa: %d\n"+
					"Essas foram as suas tentativas: %v\n",
				x,          // preenche o primeiro %d
				i+1,        // preenche o segundo %d
				chutes[:i], // preenche o %v
			)
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Infelizmente, você não acertou o número, que era %d\n"+
			"Essas foram as suas tentativas: %v\n",
		x, chutes,
	)
}
