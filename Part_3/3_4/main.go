package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	Bat string
}

func (b Battery) String() string {
	var count int
	for _, v := range b.Bat {
		if v == '0' {
			count++
		}
	}
	return fmt.Sprintf("[%s%s]", strings.Repeat(" ", count), strings.Repeat("X", 10-count))
}

func main() {
	//var str string
	//fmt.Scan(&str)
	//batteryForTest := Battery{
	//	Bat: str,
	//}
}
