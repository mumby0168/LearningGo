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

func usingRange() {
	fmt.Println("range example")
	var data []int
	data = append(data, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i, index := range data {
		fmt.Printf(" %v * 2 = %v\n", index, i*2)
	}

}

func omitRange() {
	fmt.Println("range example with omitted options")
	var data []int
	data = append(data, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	for _, value := range data {
		fmt.Println(value)
	}

	for index := range data {
		fmt.Println(index)
	}

}

func main() {
	fmt.Println("Standard for loop")
	standardForLoop()

	fmt.Println("Standard while loop")
	whileLoop()
	usingRange()
	omitRange()
}
