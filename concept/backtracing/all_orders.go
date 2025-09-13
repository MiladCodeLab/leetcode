package main

import "fmt"

func allOrders(input []byte, l int) [][]byte {
	var result [][]byte

	if len(input) == l {
		c := make([]byte, len(input))
		copy(c, input)
		return [][]byte{c}
	}
	for i := l; i < len(input); i++ {
		input[l], input[i] = input[i], input[l]
		//result = append(result, input)
		result = append(result, allOrders(input, l+1)...)
		input[l], input[i] = input[i], input[l]

	}
	//result = append(result, allOrders(input, l+1)...)

	return result
}

func main() {
	println("all orders")
	res := allOrders([]byte{1, 2, 3, 4}, 0)
	fmt.Printf("result: %+v", res)
}
