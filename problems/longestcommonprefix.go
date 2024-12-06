package problems

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Find the minimum string length
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}

	// Check characters one by one
	for charIndex := 0; charIndex < minLength; charIndex++ {
		for _, str := range strs {
			if str[charIndex] != strs[0][charIndex] {
				return str[:charIndex]
			}
		}
	}

	return strs[0][:minLength]
}
