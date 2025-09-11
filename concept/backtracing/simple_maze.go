package main

func walk(start, target int) {
	if start == target {
		println("found the target ", target)
		return
	}

	println("arrived at step ", start)

	walk(start+1, target)

	println("leaving at step ", start)
}

func main() {
	println("simple maze")
	walk(1, 5)
}
