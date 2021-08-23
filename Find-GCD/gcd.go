package gcd

// findGCD returns Greatest Common Divisor (GCD) of given slice of numbers
func FindGCD(nums ...int) int {
	var p, gcd int

	// Generator for prime numbers
	closeCh := make(chan int)
	defer close(closeCh)
	ch := primeGenerator(closeCh)

	gcd = 1
	p = <-ch
	for allGreaterThanOrEqual(nums, p) {
		// When all numbers are greater than prime number p
		for allPerfectlyDivisible(nums, p) {
			// When all numbers are perfectly divisible by prime number p
			// Iterate to see if same prime number can divide multiple times
			gcd *= p

			// Divide prime number p from all the numbers
			nums = divideFromAll(nums, p)
		}

		// Read higher prime number for next iteration
		p = <-ch
	}
	return gcd
}
