// Custom types
package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelicius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294
	c := kelvinToCelicius(k)
	fmt.Print(k, " K is ", c, "c")
}
