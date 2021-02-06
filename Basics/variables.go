package main 

var c, python, java bool // Vars claimed at the function level
var i, j int = 1, 2  // vars with initialization
const Pi = 3.14 // Cannot use the := 

func variables() {

	func printPackageLevelVariables() {
		var i int // var claimed within func scope
		fmt.Println(i, c, python, java)
	}

	func printInitialzedVariables() {
		fmt.Println(i, j)
	}

	func printImplicitlyTypedVars () {
		implicit1 := "hi"
		implicit2, implicit3 := 1, 2

		fmt.Println(implicit1, implicit2, implicit3);
	}

	func printConstant() {
		fmt.Println(Pi)
	}
	

}