package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Empty")
	}
	r := csv.NewReader(f)
	record, _ := r.ReadAll()
	if len(record) == 10 {
		fmt.Println(record[4][2])
		fmt.Printf(path)
	}

	return nil
}

func main() {
	const root = "."

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}
