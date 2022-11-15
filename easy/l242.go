package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var S = make(map[byte]int, len(s))
	var T = make(map[byte]int, len(t))

	for i := 0; i < len(s); i++ {
		S[s[i]]++
		T[t[i]]++
	}

	for k, v := range S {
		if v != T[k] {
			return false
		}
	}

	for k, v := range T {
		if v != S[k] {
			return false
		}
	}

	return true
}
