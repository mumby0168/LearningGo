package main

import "fmt"

func basicArrays() [6]int {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	return primes
}

func basicSlices() {
	primes := basicArrays()
	firstThree := primes[0:3]
	fmt.Println(firstThree)
}

func referenceSlices() {
	primes := basicArrays()
	slice := primes[0:3]
	//also changes value in array
	slice[0] = 25
	fmt.Println(primes)
}

func sliceLiterals() {
	//dynamically sized unlike array
	slice := []bool{true, false, false}
	fmt.Println(slice)
}

func sliceUsingMake() {
	a := make([]int, 5) // zeros out enough elements
	fmt.Println(len(a))
	b := make([]int, 5, 10) // second args gives capacity
	fmt.Println(len(b), cap(b))
}

func appendingToASlice() {
	var a []int
	a = append(a, 1, 2, 4, 5, 6, 7, 8, 9)
	fmt.Println(a)

}

func main() {
	basicArrays()
	basicSlices()
	referenceSlices()
	sliceLiterals()
	sliceUsingMake()
	appendingToASlice()
}
