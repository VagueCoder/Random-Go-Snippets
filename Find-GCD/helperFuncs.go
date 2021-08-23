package gcd

// allGreaterThanOrEqual iterates over slice to check all elements are greater than given number
func allGreaterThanOrEqual(arr []int, n int) bool {
	for _, a := range arr {
		if a < n {
			return false
		}
	}
	return true
}

// allPerfectlyDivisible iterates over slice to check all elements are divisible by given number
func allPerfectlyDivisible(arr []int, n int) bool {
	for _, a := range arr {
		if a%n != 0 {
			return false
		}
	}
	return true
}

// divideFromAll iterates over slice to divide all elements by given number
func divideFromAll(arr []int, n int) []int {
	var result []int
	for _, a := range arr {
		a /= n
		result = append(result, a)
	}
	return result
}
