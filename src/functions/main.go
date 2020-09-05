package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (x2, y2 int) {
	y2, x2 = y, x
	return
}

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(swap(10, 20))
	x2, y2 := swap(10, 20)
	fmt.Println(x2, y2)
}