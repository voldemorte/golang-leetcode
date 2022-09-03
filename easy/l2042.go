package easy

import (
	"strings"
)

func areNumbersAscending(s string) bool {
	// Split the given string into an array of words
	words := strings.Fields(s)

	var numbers []int
	var matchNumbers = map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	for _, word := range words {
		// Check if the word is a number
		if word[0] >= '0' && word[0] <= '9' {
			// Convert the string numeral into int.
			temp := 0
			for i := 0; i < len(word); i++ {
				temp = (temp * 10) + matchNumbers[rune(word[i])]
			}

			if len(numbers) != 0 && temp <= numbers[len(numbers)-1] {
				// Check if the new numeral meets the requirement of the problem
				return false
			}
			numbers = append(numbers, temp)
		}
	}

	if len(numbers) == 0 {
		return false
	}

	return true
}
