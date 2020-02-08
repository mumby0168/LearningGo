package main

import "fmt"

/*
Works very much like C/C++ pointers
NO POINTER ARITHMETIC
*/

func basic() {
	number := 10
	pNumber := &number

	fmt.Println(pNumber)  // prints mem address
	fmt.Println(number)   // prints value
	fmt.Println(*pNumber) // prints value

}

func main() {
	basic()
}
