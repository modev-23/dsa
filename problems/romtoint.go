package problems

func romanToInt(s string) int {
	result := 0

	intMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s)-1; i++ {
		if intMap[s[i]] >= intMap[s[i+1]] {
			result += intMap[s[i]]
			continue
		}
		result -= intMap[s[i]]
	}
	result += intMap[s[len(s)-1]]
	return result
}
