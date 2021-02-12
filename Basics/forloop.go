package main

import (
	"fmt"
	"math"
)


func forloops() {
	
	func forLoopExample() {
		sum := 0

		for i := 0; i < 10; i++ {
			sum += i 
		}

		fmt.Println(sum)
	}

	func shortForLoop() {
		sum := 1

		for ; sum < 1000; {
			sum += sum 
		}

		fmt.Println(sum)
	}

	func forWhileLoop() {
		sum := 1

		for sum < 1000 {
			sum += sum
		}

		fmt.Println(sum)
	}

	func foreverForLoop() {
		for {
		}
	}

}