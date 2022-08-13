package main

func main() {
	vote(0, 0, 1)
}
func vote(x int, y int, z int) int {
	sum := x + y + z
	if sum < 2 {
		return 0
	} else {
		return 1
	}
}
