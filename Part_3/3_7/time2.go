package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	gg := []byte(reader)
	ggStr := gg[:len(gg)-1]
	myTime, err := time.Parse("2006-01-02 15:04:05", string(ggStr))
	if err != nil {
		fmt.Println(err)
	}
	if myTime.Hour() < 13 {
		fmt.Println(myTime.Format("2006-01-02 15:04:05"))
	} else {
		//future := myTime.Add(time.Hour * 24)
		//fmt.Println(future.Format("2006-01-02 15:04:05"))
		future := myTime.AddDate(0, 0, 1)
		fmt.Print(future.Format("2006-01-02 15:04:05"))
	}
}
