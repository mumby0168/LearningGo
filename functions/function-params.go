package main

import "fmt"

//takes a func as a param which takes two args and returns a float
func compute(execute func(float64, float64) float64) float64 {
	return execute(3, 4)
}

func main() {
	value := compute(func(x, y float64) float64 {
		return x * y
	})

	fmt.Println(value)
}
