package main

import (
	"fmt"
	"strconv"
)

// bool
// int int8 int16 int32 int64
// uint uint8 uint32 uint64 uintptr
// byte = uint8
// rune = int32
// float 32, float 64
// complex64 complex128
// string

func main() {
	var b uint8 = 10
	takeByte(b)
	var x int = 10084
	f := float64(x)
	fmt.Println(f)
	//s := string(x)
	//fmt.Printf(s)

	s := strconv.FormatInt(int64(x), 10) // 10 é a base

	fmt.Println(s)

	const y = 10
	takeint64(y)

}

func takeByte(b byte) {
	fmt.Println(b)
}

func takeint64(x int64) {
	fmt.Println(x)
}
