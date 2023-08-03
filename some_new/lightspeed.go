package main

import "fmt"

// it all starts here
func main() {
	const lightSpeed = 232424
	var distance = 50000
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 575800
	fmt.Println(distance/lightSpeed, "seconds")
}
