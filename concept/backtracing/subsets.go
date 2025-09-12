package main

import "fmt"

func subsets(input []byte) [][]byte {
	var result [][]byte
	if len(input) == 0 {
		return [][]byte{}
	}
	temp := []byte{}
	for i := 0; i < len(input); i++ {
		temp = append(temp, input[i])
		result = append(result, temp)
	}

	result = append(result, subsets(input[1:])...)

	return result
}

func main() {
	println("subsets")
	res := subsets([]byte{1, 2, 3, 4})
	fmt.Printf("result: %+v", res)
}
