// Optional patterns using optional functions
package main

import "fmt"

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

// `NewHouse` is a constructor for `*House`
func NewHouse(opts ...HouseOption) *House {
	const (
		defaultFloors    = 2
		defaultFireplace = true
		defaultMaterial  = "wood"
	)
	h := &House{
		Material:     defaultMaterial,
		HasFireplace: defaultFireplace,
		Floors:       defaultFloors,
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *House as the argument
		opt(h)
	}

	return h
}

type HouseOption func(*House)

func WithConcrete() HouseOption {
	return func(h *House) {
		h.Material = "Concrete"
	}
}

func WithoutFireplace() HouseOption {
	return func(h *House) {
		h.HasFireplace = false
	}
}

func WithFloors(floors int) HouseOption {
	return func(h *House) {
		h.Floors = floors
	}
}

func main() {
	h := NewHouse(
		WithConcrete(),
		WithoutFireplace(),
	)

	fmt.Printf("%+v", h)
	// &{Material:concrete HasFireplace:false Floors:2}
}
