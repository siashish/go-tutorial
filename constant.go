package main

import "fmt"

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func NeedInt(x int) int           { return x*20 + 1 }
func NeedFloat(x float64) float64 { return x * 0.1 }

func main() {
	const world = "Like and Share"
	fmt.Println("Please", world)

	fmt.Println("Happy", Pi, "Day")

	const truth = true
	fmt.Println("go lang: ", truth)

	fmt.Println(NeedInt(Small))
	fmt.Println(NeedFloat(Small))
	fmt.Println(NeedFloat(Big))
}
