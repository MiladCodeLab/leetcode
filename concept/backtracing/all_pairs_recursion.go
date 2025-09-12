package main

import "fmt"

func allPairs(input []byte) [][]byte {
	var result [][]byte

	if len(input) < 2 { // base case
		return [][]byte{}
	}

	temp := []byte{}
	for i := 1; i < len(input); i++ {
		temp = append(temp, input[0])
		temp = append(temp, input[i])
		result = append(result, temp)
		temp = []byte{}
		// result = append(result, []byte{input[0], input[i]})
	}

	res := allPairs(input[1:])
	for _, v := range res {
		result = append(result, v)
	}
	// result = append(result, allPairs(input[1:])...)
	// fmt.Printf("data: %+v\n", result)
	return result
}

func main() {
	println("all pairs recursion")
	res := allPairs([]byte{1, 2, 3, 4})
	fmt.Printf("result: %+v", res)
}
