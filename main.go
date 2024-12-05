package main

import (
	"fmt"

	"modev.dsa/problems"
)

func main() {
	w1, w2 := "hello", "world1"

	result := problems.MergeAlternately(w1, w2)

	fmt.Print(result)
}
