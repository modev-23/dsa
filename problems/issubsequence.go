package problems

func isSubsequence(s string, t string) bool {
	i, j := 0, 0 // Two pointers: i for s, j for t

	// Traverse t while looking for characters in s
	for j < len(t) {
		// If characters match, move i to the next character in s
		if i < len(s) && s[i] == t[j] {
			i++
		}
		// Always move j to the next character in t
		j++
	}

	// If we've matched all characters in s, i will reach len(s)
	return i == len(s)
}
