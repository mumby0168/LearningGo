package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func initilization() {
	//almost a ctor
	v1 := vertex{1, 2}
	v2 := vertex{X: 1}
	fmt.Println(v1)
	fmt.Println(v2)
}

func main() {
	point := vertex{1, 2}
	pPoint := &point
	fmt.Println(point.X, point.Y)
	//accessing a pointer or a direct object is the same (.)
	fmt.Println(pPoint.X, pPoint.Y)
	initilization()
}
