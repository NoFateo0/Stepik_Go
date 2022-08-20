package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	gg := strings.Split(reader, ",")
	//13.03.2018 14:00:15,12.03.2018 14:00:15
	myTime, err := time.Parse("02.01.2006 15:04:05", gg[0])
	if err != nil {
		fmt.Println(err)
	}
	gg[1] = strings.TrimSpace(gg[1])
	myTime2, err := time.Parse("02.01.2006 15:04:05", gg[1])
	if err != nil {
		fmt.Println(err)
	}
	if myTime2.After(myTime) {
		res := time.Since(myTime2) - time.Since(myTime)
		fmt.Println(res.Round(time.Second))
	} else {
		res := time.Since(myTime) - time.Since(myTime2)
		fmt.Println(res.Round(time.Second))
	}
}
