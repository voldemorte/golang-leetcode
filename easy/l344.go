package easy

func reverseString(s []byte) {
	var temp byte
	for i := 0; i < len(s)/2; i++ {
		temp = s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = temp
	}
}
