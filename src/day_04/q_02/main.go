package main

import "fmt"

func main() {
	start := 254032
	end := 789860

	test1 := 112233
	test2 := 123444
	test3 := 111122
	test4 := 117722
	test5 := 111111
	test6 := 111224
	fmt.Printf("test 1: %d -> %v\n\n", test1, IsValidPassword(test1))
	fmt.Printf("test 2: %d -> %v\n\n", test2, IsValidPassword(test2))
	fmt.Printf("test 3: %d -> %v\n\n", test3, IsValidPassword(test3))
	fmt.Printf("test 4: %d -> %v\n\n", test4, IsValidPassword(test4))
	fmt.Printf("test 5: %d -> %v\n\n", test5, IsValidPassword(test5))
	fmt.Printf("test 6: %d -> %v\n\n", test6, IsValidPassword(test6))

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
	curr := int(input / 100000)
	hasDouble := false
	digitCount := 1 // count the first digit right away
	for i := 10000; i > 0; i /= 10 {
		nextDigit := int(input/i) % 10
		// fmt.Printf("> %d - %d\n", curr, nextDigit)
		if curr > nextDigit {
			return false
		}

		if curr == nextDigit {
			digitCount++
		} else { // different digit. check to see if the digit count was exactly 2
			if !hasDouble && digitCount == 2 {
				// fmt.Println("found Double")
				hasDouble = true
			}

			digitCount = 1
		}

		curr = nextDigit
	}

	// it's possible the last two digits make the "double". Check it:
	if !hasDouble && digitCount == 2 {
		// fmt.Println("found Double")
		hasDouble = true
	}

	// if we got through the loop above, then the digits are in ascending order.
	// just return if the double was present
	return hasDouble
}
