package main

import (
	"fmt"
)

func reverseNum(num int) int {
	rev := 0
	for num > 0 {
		rev = (rev * 10) + (num % 10)
		num = num / 10
	}

	return rev

}

func main() {
	num := 8668
	answer := "is not"

	if num == reverseNum(num) {
		answer = "is"
	}

	fmt.Println(fmt.Sprintf("The number, %d, %s a palindrome number", num, answer))
}
