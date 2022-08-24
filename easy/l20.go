package easy

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {

	// map data to compare the strings
	var parenthesis = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	// Provision a stack of max capacity to the length of input string
	stack := make([]rune, 0, len(s))

	// Convert the string to list of runes
	input := []rune(s)

	// Iterate over the input
	for _, v := range input {

		// Get the head of stack
		lastIndex := len(stack) - 1

		// stack is empty. Push the rune in
		// This is the default condition but added to address invalid index error.
		if lastIndex < 0 {
			stack = append(stack, v)
		} else if v == parenthesis[stack[lastIndex]] { // Check if the character matches a closing brace
			// Pop the last pushed character
			stack = stack[:lastIndex]
		} else {
			// Default action to push the rune into stack
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}
