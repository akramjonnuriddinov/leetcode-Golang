package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	max := 0
	for i := len(prices) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > max {
				max = prices[i] - prices[j]
			}
		}
	}

	return max
}
