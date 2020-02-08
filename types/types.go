package main

import (
	"fmt"
	"math/cmplx"
)

var (
	boolean    bool       = true
	maxInteger int64      = 1<<52 - 1
	z          complex128 = cmplx.Sqrt(10)
)

func typeConversion() {
	i := 52
	f := float64(i)
	u := uint(f)

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)
}

func constants() {
	const i = 10
	const world = "World"

	fmt.Println(i)
	fmt.Println(world)
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", boolean, boolean)
	fmt.Printf("Type: %T Value: %v\n", maxInteger, maxInteger)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	typeConversion()
}
