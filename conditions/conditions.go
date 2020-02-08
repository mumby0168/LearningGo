package main

import (
	"fmt"
	"math"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func inlineIf(x int) bool {
	if a := x + 10; a > 20 {
		return true
	}
	return false
}

func inlineIfElse(x int) {
	if a := x % 2; a != 0 {
		fmt.Printf("%v is not a mutliple of 2\n", x)
	} else {
		fmt.Printf("%v is a mutliple of 2\n", x)
	}
}

//no breaks or returns required
func basicSwitch(statement, name string) {
	switch statement {
	case "Hello":
		fmt.Println("Hello " + name)
	case "Goodbye":
		fmt.Println("Goodbye " + name)
	case "Morning":
		fmt.Println("Good morning, " + name)
	case "Afternoon":
		fmt.Println("Good afternoon, " + name)
	default:
		fmt.Println("Please choose a register greeting")
	}
}

func boolSwitch() {
	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("Good afternoon")
	case t.Hour() < 12:
		fmt.Println("Good morning")
	}
}

func main() {
	fmt.Println(sqrt(4), sqrt(-4))
	fmt.Println(inlineIf(3), inlineIf(15))
	inlineIfElse(22)
	inlineIfElse(25)
	basicSwitch("Hello", "Billy")
	basicSwitch("Afternoon", "Billy")
	boolSwitch()
}
