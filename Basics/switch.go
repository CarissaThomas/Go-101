package main

import (
	"fmt"
	"runtime"
)


func switchStatements() { 
	
	func switchStatement() {
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("OS X")
		}
	}

}