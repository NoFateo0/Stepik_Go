package main

import (
	"fmt"
	"time"
)

func main() {
	var strTime string
	fmt.Scan(&strTime)

	formTime, _ := time.Parse(time.RFC3339, strTime)
	fmt.Println(formTime.Format(time.UnixDate))
}
