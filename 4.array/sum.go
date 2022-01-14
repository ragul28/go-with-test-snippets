package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numberToSum ...[]int) []int {
	var sums []int

	for _, numers := range numberToSum {
		sums = append(sums, Sum(numers))
	}

	return sums
}

func SumAllTails(numberToSum ...[]int) []int {
	var sums []int

	for _, numers := range numberToSum {

		if len(numers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
