package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}