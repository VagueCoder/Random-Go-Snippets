package gcd

// PrimeGenerator creates a generator goroutine that returns prime numbers, least first
func primeGenerator(close <-chan int) chan int {
	var current int
	var flag bool
	ch := make(chan int)

	go func() {
		current = 1
		for {
			select {
			case <-close:
				// When channel is closed, i.e., requested for closure
				return

			default:
				if current < 3 {
					// For prime numbers 2 & 3
					current++
					ch <- current
				} else {
					// For Prime Numbers greater than 3
					for i := current + 1; ; i++ {
						flag = true
						for j := 2; j <= i/2; j++ {
							if i%j == 0 {
								flag = false
								break
							}
						}

						// When the number is not perfectly divisible by any number except 1 & itself.
						if flag {
							ch <- i
							current = i
							break
						}
					}
				}
			}
		}
	}()

	return ch
}
