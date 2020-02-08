package main

import "fmt"

func standardForLoop() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func infinite() {
	for {

	}
}

func main() {
	fmt.Println("Standard for loop")
	standardForLoop()

	fmt.Println("Standard while loop")
	whileLoop()

}
