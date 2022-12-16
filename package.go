package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(4, 5))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(12))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
