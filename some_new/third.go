// Floating precision and formating

package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("Third: %2f\n", third)
	fmt.Printf("Third: %05.3f\n", third)
	fmt.Printf("Third: %0.2f", third)
}
