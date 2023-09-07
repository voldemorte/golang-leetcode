package main

import (
	"fmt"
)

func numFactorial(num int) int {
	if num == 1 {
		return num
	}
	return numFactorial(num-1) * num
}

func main() {
	num := 6

	fmt.Println(fmt.Sprintf("The factorial of %d is %d", num, numFactorial(num)))
}
