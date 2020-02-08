package main

import (
	"fmt"
	"os"
)

func forLoopEcho() {

	var s, seperator string

	for i := 1; i < len(os.Args); i++ {
		s += seperator + os.Args[i]
		s = " "
	}

	fmt.Println("Simple for loop to pring args:")
	fmt.Println(s)
}

func forEachLoopEcho() {
	var s, seperator string

	for _, argument := range os.Args[1:] {
		s += seperator + argument
		seperator = " "
	}

	fmt.Println("Kind of foreach loop using _ to ignore the index")
	fmt.Println(s)
}

func main() {
	forEachLoopEcho()
}
