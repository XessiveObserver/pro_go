// Own types
package main

import "fmt"

func main() {
	type celicius float64
	const degrees = 20

	var temperature celicius = degrees

	temperature += 10

	fmt.Println(temperature)
}
