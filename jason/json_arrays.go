// Json Arrays
package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	birdJson := `[
		{"species": "pigeon", "description": "perchs on rocks"},
		{"species": "eagle", "description": "bird of prey"}
	]`

	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)

	fmt.Println("Birds:", birds)
}
