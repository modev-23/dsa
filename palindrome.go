package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
