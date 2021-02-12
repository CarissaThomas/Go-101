package main

import (
	"fmt"
	"math"
)

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32 & represents a Unicode code point
float32 float64
complex64 complex128

func types() {

	func printZeroValues() {
		var i int // 0
		var f float64 // 0 
		var b bool // false 
		var s string // ""

		fmt.Printf(i, f, b, s);
	}

	func printTypeConversions() {
		var x, y int = 3, 4 
		var f float64 = math.Sqrt(float64(x*x + y*y))
		var z uint = uint(f)
		fmt.Println(x, y, z)
	}

	func printConstant() {
		const test = "This is a test"
		fmt.Println(test);

		const isTest = true;
		fmt.Println(isTest);
	}

	
}