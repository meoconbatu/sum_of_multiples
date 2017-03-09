package summultiples

import "sort"

func SumMultiples(limit int, divisors ...int) (sum int) {
	var multiples []int
	for _, divisor := range divisors {
		multiples = append(multiples, GetMultiplesUpTo(divisor, limit)...)
	}
	sort.Ints(multiples)
	numMultiples := len(multiples)
	for i := 0; i < numMultiples; i++ {
		if i+1 == numMultiples {
			return sum + multiples[i]
		}
		if multiples[i] != multiples[i+1] {
			sum += multiples[i]
		}
	}
	return
}

func GetMultiplesUpTo(divisor int, limit int) (multiples []int) {
	for i := 0; i*divisor < limit; i++ {
		multiples = append(multiples, i*divisor)
	}
	return
}
