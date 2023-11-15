package wheel

import (
	"math/rand"
)

type Wheel struct {
	Num_of_cells int8
	Slots        Slot
}

func (w Wheel) Spin() float64 {
	weights := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}

	sum := 0.0
	for _, weight := range weights {
		sum += weight
	}

	value := rand.Float64() * sum

	for _, weight := range weights {
		value -= weight
		if value <= 0 {
			return weight
		}
	}

	return weights[len(weights)-1]
}

type Slot struct {
	Name   string
	Image  string
	Chance int8
}
