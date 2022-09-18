package interviewquestions

import "fmt"

// The code will take two strings as input in two lines. The first string will contain letters on the crossword("cwd"), and the second string will contain a word that you want to check.(i.e. you want to check whether the word can be made out of given letters in the same order). The output should say "yes" if the word can be formed, else it should say "no"
// Sample Input 1: ccwd crossword
// Sample Output1: no
// Sample Input 2: cwd crossword
// Sample Output 2: yes
// In the interest of time hard-code the the strings into variables
func main() {
	s1 := "cwd"
	s2 := "crossword"

	var index int

	for i := 0; i < len(s2); i++ {
		if index < len(s1) && s1[index] == s2[i] {
			index++
		}
	}

	if index == len(s1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
