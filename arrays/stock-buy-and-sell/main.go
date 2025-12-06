package main

import "fmt"

func StockBuyAndSell(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		profit := price - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(StockBuyAndSell([]int{7, 10, 1, 3, 6, 9, 2})) 
	fmt.Println(StockBuyAndSell([]int{7, 6, 4, 3, 1}))        
	fmt.Println(StockBuyAndSell([]int{1, 3, 6, 9, 11}))       
}
