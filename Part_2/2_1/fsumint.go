package main

func main() {
	sumInt(0, 1)
}

func sumInt(a ...int) (int, int) {
	var sum int

	for _, v := range a {
		sum += v
	}
	return len(a), sum
}
