package main

import (
	"fmt"
	"math"
)


func ifStatements() {

	func ifStatement(x float64) {
		if x < 0 {
			return x;
		}
	}

	func ifStatementWithShortStatement(lim float64) {
		if v:= 1; v < lim {
			return v
		} else {
			fmt.Println(lim)
		}
		return lim
	}



}