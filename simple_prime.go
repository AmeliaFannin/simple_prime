package simple_prime

func PrimeGenerator(max int) chan int {
	out := make(chan int)
	primes := []int{2}

	go func() {
		out <- 2
		for i := 3; i <= max; i += 2 {
			if primeCheck(primes, i) {
				out <- i
				primes = append(primes, i)
			}
		}
		close(out)
	}()
	return out
}

func primeCheck(primes []int, candidate int) bool {
	for _, known_prime := range primes {
		if candidate%known_prime == 0 {
			return false
		}
	}
	return true
}
