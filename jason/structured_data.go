// Marshalling and Unmarshalling json using go

package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string
	Sound string
}

func main() {
	animalJson := `{"name": "cow", "sound": "mows"}`
	var animal Animal
	json.Unmarshal([]byte(animalJson), &animal)
	fmt.Printf("Name: %v, Sound: %v", animal.Name, animal.Sound)
}
