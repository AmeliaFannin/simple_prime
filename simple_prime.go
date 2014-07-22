package simple_prime

// sends possible primes, less than max, to primeCheck function
func PrimeGenerator(max int) chan int {
	out := make(chan int)
	primes := []int{2}

	go func() {
		out <- 2
		for i := 3; i <= max; i += 2 {
			if primeCheck(primes, i) {

				// puts confirmed primes into the out channel
				out <- i

				// adds confirmed primes to the array,
				primes = append(primes, i)
			}
		}
		close(out)
	}()
	// returns channel of confirmed primes
	return out
}

// checks possible prime for divisibility by confirmed primes
func primeCheck(primes []int, candidate int) bool {
	for _, known_prime := range primes {
		if candidate%known_prime == 0 {
			return false
		}
	}
	return true
}
