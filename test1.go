package main

import (
	"fmt"
)

func plus(a int64, b int64) int64 {
	return a + b
}

func main() {
	var var1 int64
	var var2 int64

	fmt.Print("1: ")
	fmt.Scan(&var1)

	fmt.Print("2: ")
	fmt.Scan(&var2)

	fmt.Println("result:", plus(var1, var2))
}
