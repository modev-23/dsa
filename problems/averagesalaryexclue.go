package problems

func Average(salary []int) float64 {

	result, min, max := 0, salary[0], salary[0]

	for i := range salary {
		result += salary[i]

		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
	}

	result = (result - min - max) / (len(salary) - 2)

	return float64(result)

}
