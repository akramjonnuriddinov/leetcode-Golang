package main

import (
	"errors"
	"fmt"
)

func main2() {
	result, err := divide(10, 9)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res: ", result)
}

func divide(a, b float64) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return int(a / b), nil
}
