package main

import "fmt"

func allPairs(input []byte) [][]byte {
	var result [][]byte

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			result = append(result, []byte{input[i], input[j]})
		}
	}
	return result
}

func main() {
	println("all pairs")
	res := allPairs([]byte{'A', 'B', 'C', 'D'})
	fmt.Printf("result: %+v", res)
}
