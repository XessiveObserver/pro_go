// Method chaining
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) withName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person) withAge(age int) *Person {
	p.Age = age
	return p
}

func main() {
	p := &Person{}

	p = p.withName("John").withAge(21)

	fmt.Println(*p)
}
