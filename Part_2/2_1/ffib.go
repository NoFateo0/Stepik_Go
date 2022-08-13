package main

func main() {
	fibonacci(4)
}

func fibonacci(n int) int {
	var a, b int

	a = 1
	b = 1
	i := 2
	for i < n {
		a, b = b, b+a
		i++
	}
	return b
}
