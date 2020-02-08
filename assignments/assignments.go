package main

import "fmt"

func main() {
	//implicit declaration + assignment
	s := "Hello World"

	//explicit declaratiopn + assignment
	var s1 string = "Hello World"

	fmt.Println(s1)
	fmt.Println(s)
}
