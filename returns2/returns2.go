// Multiple returns in functions

package main

import (
	"fmt"
)

func fullName() (firstName string, secondName string) {
	firstName = "fools"
	secondName = "gold"
	return
}

func main() {
	name1, name2 := fullName()
	fmt.Printf("%v %v", name1, name2)
}
