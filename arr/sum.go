package arr

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(numbersArr ...[]int) []int {
	var sums []int
	for _, numbers := range numbersArr {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersArr ...[]int) []int {
	var sums []int
	for _, numbers := range numbersArr {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
