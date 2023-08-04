// Struct Embedding
package main

import "fmt"

type People struct {
	Name   string
	Age    int
	Height float64
}

type Person struct {
	People // Has features of People
	Habbit string
}

func (p Person) Move() {
	fmt.Printf("Name %v, Age %v, Height %.2f, Habbit %v", p.Name, p.Age, p.Height, p.Habbit)
}

func main() {
	person := Person{People: People{Name: "Okello", Age: 35, Height: 6.9}, Habbit: "cycling"}
	person.Move()
}
