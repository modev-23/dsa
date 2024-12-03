package problems

func MergeAlternately(word1, word2 string) string {
	a, b := 0, 0
	A, B := len(word1), len(word2)
	result := make([]byte, 0, A+B) // Preallocate memory for efficiency

	for a < A && b < B {
		result = append(result, word1[a])
		a++
		result = append(result, word2[b])
		b++
	}

	for a < A {
		result = append(result, word1[a])
		a++
	}

	for b < B {
		result = append(result, word2[b])
		b++
	}

	return string(result)
}
