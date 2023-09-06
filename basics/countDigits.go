package main

import (
	"fmt"
)

// This function simply counts the number of times itself was called.
func countDigit(num int) int {
	if num < 10 {
		return 1
	}
	return countDigit(num/10) + 1
}

func main() {
	num := 74658

	fmt.Println(fmt.Sprintf("the number of digits in %d is %d", num, countDigit(num)))
}
