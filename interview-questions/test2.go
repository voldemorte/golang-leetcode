package interviewquestions

// Day 1 — Evaluate a String Expression (no parentheses)
//
// Problem:
// Given a string s representing a non-negative integer arithmetic expression containing digits and the operators +, -, *, /, evaluate it and return the integer result.
//
// There are no parentheses.
//
// Division is integer division (truncate toward zero).
//
// Input may contain spaces; ignore them.
//
// Examples
// Input:  "3+2*2"
// Output: 7
// Explanation: 3 + (2*2) = 7
//
// Input:  " 14-3/2 "
// Output: 13
// Explanation: 14 - (3/2) = 14 - 1 = 13
//
// Input:  "100/10*3+6-4/3"
// Output: 35
// Explanation: ((100/10)*3) + 6 - (4/3) = (10*3) + 6 - 1 = 35
//
// Constraints
//
// 1 <= len(s) <= 10^5
//
// s contains digits, + - * /, and spaces only.
//
// The expression is valid (no leading operator, no invalid sequences).
//
// Goal
//
// Implement func calculate(s string) int in Go.
//
// Aim for O(n) time, O(1) or O(n) extra space (both acceptable).
//
// Edge Cases to handle
//
// Multiple spaces between tokens: " 3 + 5 / 2 "
//
// Long numbers (not just single digits): "123*45+6"
//

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	res := 0
	last := 0
	num := 0

	op := byte('+')

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if unicode.IsDigit(rune(ch)) {
			num = num*10 + int(ch-'0')
		}

		if (!unicode.IsDigit(rune(ch)) && ch != ' ') || i == len(s)-1 {
			switch op {
			case '+':
				res += last
				last = num
			case '-':
				res += last
				last = -num
			case '*':
				last = last * num
			case '/':
				last = last / num
			}
			op = ch
			num = 0
		}
	}

	return res + last
}

func main() {
	tests := []struct {
		in   string
		want int
	}{
		{"3+2*2", 7},
		{"3*2+2", 8},
		{" 14-3/2 ", 13},
		{"100/10*3+6-4/3", 35},
		{"  3   + 5 /  2 ", 5},
		{"123*45+6", 5541},
		{"7-10/3", 4}, // 7 - 3 = 4 (trunc toward zero)
	}
	for _, tc := range tests {
		got := calculate(tc.in)
		fmt.Printf("input=%q -> got=%d (want=%d)\n", tc.in, got, tc.want)
	}
}
