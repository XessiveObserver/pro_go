package main

import "fmt"

func arith(a int, b int) (int, int) {
	add := a + b
	sub := b - a

	return add, sub
}

func main() {
	fmt.Println(arith(3, 4))
	fmt.Println(arith(3, 8))
}
