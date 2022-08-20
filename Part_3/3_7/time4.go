package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//12 мин. 13 сек.
	reader = strings.TrimSpace(reader)
	reader = strings.ReplaceAll(reader, " ", "")
	reader = strings.ReplaceAll(reader, "мин.", "m")
	reader = strings.ReplaceAll(reader, "сек.", "s")
	gg, err := time.ParseDuration(reader)
	if err != nil {
		fmt.Println(err)
	}
	myCon := time.Unix(now, 0).UTC()
	lol := myCon.Add(gg)

	fmt.Println(lol.Format(time.UnixDate))
}
