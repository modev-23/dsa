package main

import (
	"fmt"

	"modev.dsa/problems"
)

func main() {
	salary := []int{100, 200, 600, 400, 500}
	r := problems.Average(salary)

	fmt.Println(float64(r))
}
