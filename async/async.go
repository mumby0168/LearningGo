package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin")

	go func() { fmt.Println("Hello World") }()
	time.Sleep(1 * time.Second)

	fmt.Println("end")
}
