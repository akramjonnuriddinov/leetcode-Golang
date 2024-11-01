package main

import "fmt"

func main() {
	fmt.Println(addDigits(1234))
}

func addDigits(num int) int {
	sum := 0
	num = num / 10
	for num > 0 {
		sum += num
	}
	return sum
}
