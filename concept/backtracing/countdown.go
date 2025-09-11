package main

func countdown(n int) {
	if n < 0 {
		return
	}
	println(n)
	countdown(n - 1)
}

func main() {
	println("recursion countdown")
	countdown(4)
}
