package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Rating []int
}

type Group struct {
	Students []Student
}

type Rating struct {
	Average float32
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
	var g Group
	var count1 int
	var count2 int
	var av float32

	if err := json.Unmarshal(read, &g); err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(g.Students); i++ {
		count1++
		count2 += len(g.Students[i].Rating)
	}
	av = float32(count2) / float32(count1)
	rating := Rating{Average: av}
	newF, _ := json.MarshalIndent(rating, "", "    ")
	fmt.Println(string(newF))
}
