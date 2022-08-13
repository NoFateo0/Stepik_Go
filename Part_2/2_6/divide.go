package main

import "fmt"

func main() {
	result, err := divide(10, 5)
	if err != nil {
		fmt.Errorf("ошибка")
	} else {
		fmt.Println(result)
	}
}

func divide(a int, b int) (int, error) {
	var g int
	g = a / b
	return g, nil
}
