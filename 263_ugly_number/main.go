package main

import (
	"fmt"
)

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			break
		}
	}
	return n == 1
}

func main() {
	fmt.Println("ugly number")
	fmt.Printf("14 = %v\n", isUgly(14))
}
