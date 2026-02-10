package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	canal := make(chan int)

	wg.Add(1)
	go geraNum(canal)
	go verificaPar(canal, &wg)
	wg.Wait()

}

func geraNum(c chan int) {
	defer close(c)
	for {
		n := rand.Intn(100)
		c <- n
		time.Sleep(500 * time.Millisecond)
	}
}

func verificaPar(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range c {
		if n%2 == 0 {
			fmt.Printf("O número %d é par\n", n)
		} else {
			fmt.Printf("O número %d é ímpar\n", n)
		}
	}
}
