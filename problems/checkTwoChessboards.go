package problems

import (
	"fmt"
	"strconv"
	"strings"
)

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	mp := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}
	co1 := strings.Split(coordinate1, "")
	co2 := strings.Split(coordinate2, "")

	c1, _ := strconv.Atoi(co1[1])
	c2, _ := strconv.Atoi(co2[1])

	return (mp[co1[0]]+c1)%2 == (mp[co2[0]]+c2)%2
}

func main() {
	fmt.Printf("%t", checkTwoChessboards("b1", "a1"))
}
