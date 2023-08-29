package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
