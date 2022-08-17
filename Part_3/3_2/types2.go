package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var h, g string

	h = "%^80"
	g = "hhhhh20&&&&nd"

	sum := adding(h, g)
	fmt.Println(sum)
}

func adding(a, b string) int64 {
	var sum1, sum2 string

	for _, v := range a {
		if unicode.IsNumber(v) {
			sum1 += string(v)
		}
	}
	for _, v := range b {
		if unicode.IsNumber(v) {
			sum2 += string(v)
		}
	}
	aSum1, _ := strconv.Atoi(sum1)
	aSum2, _ := strconv.Atoi(sum2)

	return int64(aSum1 + aSum2)
}
