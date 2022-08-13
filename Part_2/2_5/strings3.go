package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, s string

	fmt.Scan(&x, &s)

	ind := strings.Index(x, s)
	fmt.Println(ind)
}
