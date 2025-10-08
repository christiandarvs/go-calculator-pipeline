	package main

	import "fmt"

	func AddNum(x int, y int) int {
		return x + y
	}

	func SubtractNum(x int, y int) int {
		return x - y
	}

	func MultiplyNum(x int, y int) int {
		return x * y
	}

	func DivideNum(x int, y int) int {
		return x / y
	}

	func main() {
		var x = 10
		var y = 5
		sum := AddNum(x, y)
		difference := SubtractNum(x, y)
		product := MultiplyNum(x, y)
		quotient := DivideNum(x, y)

		fmt.Printf("%d + %d = %d\n", x, y, sum)
		fmt.Printf("%d - %d = %d\n", x, y, difference)
		fmt.Printf("%d * %d = %d\n", x, y, product)
		fmt.Printf("%d / %d = %d\n", x, y, quotient)

	}
