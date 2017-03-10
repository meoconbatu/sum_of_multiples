package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	multiples := make(map[int]bool, limit)
	for _, divisor := range divisors {
		for i := divisor; i < limit; i += divisor {
			if multiples[i] == false {
				sum += i
				multiples[i] = true
			}
		}
	}
	return
}
