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

func trailingZero(num int) int {
	count := 0
	for num%10 == 0 {
		count++
		num = num / 10
	}

	return count
}

func main() {
	num := 6

	factorial := numFactorial(num)
	zeros := trailingZero(factorial)

	fmt.Println(fmt.Sprintf("The factorial of %d is %d, and has %d trailing zeros", num, factorial, zeros))
}
