package summultiples

import "runtime"

func SumMultiples(limit int, divisors ...int) (sum int) {
	numThread := runtime.NumCPU()
	shortLimitLen := limit / numThread
	if limit%numThread != 0 || shortLimitLen == 0 {
		shortLimitLen += 1
	}
	var begin, end int
	sumChan := make(chan int)
	for i := 0; i < numThread; i++ {
		begin = end + 1
		end = end + shortLimitLen
		if end  >= limit {
			end = limit - 1
		}
		go func(begin int, end int, divisors ...int) {
			sumChan <- SumMultiplesRange(begin, end, divisors...)
		}(begin, end, divisors...)
	}
	for i := 0; i < numThread; i++ {
		sum += <-sumChan
	}
	return
}

func SumMultiplesRange(begin int, end int, divisors ...int) (sum int) {
	for i := begin; i <= end; i++ {
		for _, divisor := range divisors {
			if i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return
}
