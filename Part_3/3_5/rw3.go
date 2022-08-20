package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("E:\\Sources\\Stepik_Go\\Part_3\\3_5\\test1.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var i int = 1
	for {
		line, err := reader.ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		if line == "0;" {
			fmt.Println(line, i)
		} else {
			i++
		}
	}
}
