package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var a int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		myNum, _ := strconv.Atoi(str)
		a += myNum
	}
	io.WriteString(os.Stdout, strconv.Itoa(a))
}
