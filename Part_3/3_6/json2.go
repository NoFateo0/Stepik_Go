package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type YouStruct struct {
	Id int `json:"global_id"`
}

func main() {
	file, err := os.Open("E:\\Sources\\Stepik_Go\\Part_3\\3_6\\package.json")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	read, err := io.ReadAll(reader)

	var g []YouStruct

	if err := json.Unmarshal(read, &g); err != nil {
		fmt.Println(err)
		return
	}
	var sum int
	for i := 0; i < len(g); i++ {
		sum += g[i].Id
	}
	fmt.Println(sum)
}
