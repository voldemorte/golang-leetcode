package medium

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			l := len(needle)
			if i+l <= len(haystack) && haystack[i:i+l] == needle {
				return i
			}
		}
	}
	return -1

}
