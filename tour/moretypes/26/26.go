package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	num1, num2 := 0, 0

	return func() int {
		if num2 == 0 {
			num2 = 1
			return 0
		}

		num1, num2 = num2, num1+num2
		return num1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}
