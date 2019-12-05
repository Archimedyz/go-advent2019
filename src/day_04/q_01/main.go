package main

import "fmt"

func main() {
	start := 254032
	end := 789860

	test1 := 111111
	test2 := 223450
	test3:= 123789
	fmt.Printf("test 1: %d -> %v\n", test1, IsValidPassword(test1))
	fmt.Printf("test 2: %d -> %v\n", test2, IsValidPassword(test2))
	fmt.Printf("test 3: %d -> %v\n", test3, IsValidPassword(test3))

	validPasswordCount := 0

	// brute force it
	for i := start; i <= end; i++ {
		if IsValidPassword(i) {
			validPasswordCount++
		}
	}

	fmt.Printf("count: %d\n", validPasswordCount)
}

func IsValidPassword(input int) bool {
	// first, let's check that the digits are in increasing order
	// while doing this, we can also check to see if two adjacent digits are the same!
	curr := int(input/100000)
	hasDouble := false
	for i := 10000; i > 0; i /= 10 {
		nextDigit := int(input/i) % 10
		if curr > nextDigit {
			return false
		}
		
		if curr == nextDigit {
			hasDouble = true
		}

		curr = nextDigit

	}

	// if we got through the loop above, then the digits are in ascending order.
	// just return if the double was present
	return hasDouble
}