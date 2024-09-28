// find sum of the digits of the given number without using math
package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 1234556
	fmt.Println(addDigits(number))
}

func addDigits(num int) int {
	str := numberToString(num)
	strSlice := []string{}
	for _, val := range str {
		strSlice = append(strSlice, string(val))
	}
	numSlice := stringToNumberSlice(strSlice)
	sum := 0
	for _, val := range numSlice {
		sum += val
	}
	return sum
}

func numberToString(num int) string {
	return strconv.Itoa(num)
}

func stringToNumberSlice(str []string) []int {
	res := []int{}
	for _, val := range str {
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("custom error: ", err)
		}
		res = append(res, num)
	}
	return res
}
