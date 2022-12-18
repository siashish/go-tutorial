package main

import (
	"fmt"
	"math/cmplx"
)

var (
	tobe   bool       = true
	maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 21i)
)

// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte alias of uint8
// rune alias of int32 Unicode
// float32 float64
// complex64 complex128

// zero value
// number 0
// bool flase
// string ""

func main() {
	fmt.Printf("type: %T value: %v\n", tobe, tobe)
	fmt.Printf("type: %T value: %v\n", maxint, maxint)
	fmt.Printf("type: %T value: %v\n", z, z)
}
