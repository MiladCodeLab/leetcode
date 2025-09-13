package main

import (
	"fmt"
)

func helper(c []string) []string {

	if len(c) == 1 {
		var result []string
		for _, v := range c[0] {
			result = append(result, string(v))
		}
		return result
	}
	var result []string
	var first string
	first = c[0]
	var rest []string
	rest = helper(c[1:])

	for i := 0; i < len(first); i++ {
		for j := 0; j < len(rest); j++ {
			result = append(result, string(first[i])+rest[j])
		}
	}
	return result
}

func letterCombinations(digits string) []string {
	d := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var (
		parts []string
	)

	for _, v := range digits {
		parts = append(parts, d[string(v)])
	}

	return helper(parts)
}

func main() {
	fmt.Println("letter combinations phone number")
	fmt.Printf("23 = %v\n", letterCombinations("234"))
}
