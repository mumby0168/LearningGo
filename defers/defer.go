package main

import "fmt"


// uses a stack machine and is first in last out
func stackDefered() {
	fmt.Println("Counting ...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func basicDefer() {
	//this is evaulated straight away but only executed at the return of the function
	defer fmt.Println("World")

	fmt.Println("Hello")
}

func main() {
	basicDefer()
	stackDefered()
}
