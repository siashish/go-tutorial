package main

import (
	"fmt"
	"math"
)

// value v and type T = T(v)

func main() {
	var x, y int = 4, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("Type of f :%T and value: %v\n", f, f)
	fmt.Println(x, y, z)
	var a float64
	i := a
	fmt.Printf("type of i : %T", i)
}
