package main

import (
	"fmt"
)

// ErrNegativeSqrt g
type ErrNegativeSqrt struct {
	Value float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.Value)
}

// Sqrt g
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &ErrNegativeSqrt{
			x,
		}
	}

	z := 1.0

	diff := 1.0
	for diff > 0.0001 {
		a := z
		z -= (z*z - x) / (2 * z)
		diff = a - z
		if diff < 0 {
			diff *= -1
		}
	}

	return z, nil
	// return 0, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-2))
}
