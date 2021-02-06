package main

func functions() {

	func printMe() {
		fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	}
	
	//Func takes two arguements, with a type at the end
	func twoParametersOneReturnType(x, y int) int {//return type comes after the brackets
		return x + y
	}
	
	//Takes 2 string parameters, returns 2 strings
	//calling this - a, b := multipleResults("hello", "world")
	func multipleResults(x, y string) (string, string) {
		return y, x
	}
	
	//Naked return
	func nakedReturn (sum int) (x, y int) {
		x = sum * 4/9;
		y = sum - x;
		return
	}
	
}


