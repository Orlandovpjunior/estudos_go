package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("O resultado é 2")
	default:
		fmt.Println("Algo de errado")

	}

	fmt.Println(isWeekend(time.Now()))
	fmt.Println(isWeekend2(time.Now()))
}

func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}

func isWeekend2(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return false
	default:
		return true
	}
}
