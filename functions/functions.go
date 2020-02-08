package main

import "fmt"

//omitted param types when all the same.
func add(x, y int) int {
	return x + y
}

// multiple returns from func
func swapAndConcat(a, b string) (string, string, string) {
	concat := a + b
	return b, a, concat
}

// named return values with "naked" return
func swap(a, b string) (x, y string) {
	x = b
	y = a
	return
}

func main() {

	//splits up multiple returns
	b, a, concat := swapAndConcat("Hello", "World")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(concat)

	fmt.Println(add(2, 3))
}
