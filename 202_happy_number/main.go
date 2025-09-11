package main

import (
	"fmt"
	"maps"
)

func isHappy(n int) bool {
	var computedNumber int
	seen := make(map[int]bool)
	for !seen[n] {
		seen[n] = true
		fmt.Printf("n= %d\n", n)
		for n > 0 {
			digit := n % 10
			fmt.Printf("digit= %d\n", digit)
			computedNumber += digit * digit
			n /= 10
		}
		if computedNumber == 1 {
			return true
		}

		n = computedNumber
		computedNumber = 0
	}
	if computedNumber == 1 {
		return true
	}
	for k := range maps.Keys(seen){
		fmt.Printf("number= %d\n", k)
	}

	return false
}

func main() {
	fmt.Println("happy number")
	fmt.Printf("20 = %v\n", isHappy(20))
}
