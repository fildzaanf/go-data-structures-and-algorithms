package main

import "fmt"

func MaxOneTransaction(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buy := prices[0]
	sell := 0

	for _, price := range prices {
		if price < buy {
			buy = price
		}

		if price-buy > sell {
			sell = price - buy
		}
	}

	return sell
}

func MaxTwoTransaction(prices []int) int {
	buy1 := 999999
	buy2 := 999999
	sell1 := 0
	sell2 := 0

	for _, price := range prices {

		if price < buy1 {
			buy1 = price
		}

		if price-buy1 > sell1 {
			sell1 = price - buy1
		}

		if price-sell1 < buy2 {
			buy2 = price - sell1
		}

		if price-buy2 > sell2 {
			sell2 = price - buy2
		}
	}

	return sell2
}

func MultipleTransaction(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	total := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			buy := prices[i-1]
			sell := prices[i]
			total += sell - buy
		}
	}

	return total
}

func TransactionWithFee(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	cash := 0
	stock := -prices[0]

	for _, price := range prices[1:] {

		sell := stock + price - fee
		if sell > cash {
			cash = sell
		}

		buy := cash - price
		if buy > stock {
			stock = buy
		}
	}

	return cash
}

func main() {
	fmt.Println("Max One Transaction")
	fmt.Println(MaxOneTransaction([]int{7, 10, 1, 3, 6, 9, 2})) // 8
	fmt.Println(MaxOneTransaction([]int{7, 6, 4, 3, 1}))        // 0
	fmt.Println(MaxOneTransaction([]int{1, 3, 6, 9, 11}))       // 10

	fmt.Println("\nMax Two Transactions")
	fmt.Println(MaxTwoTransaction([]int{10, 22, 5, 75, 65, 80}))      // 87
	fmt.Println(MaxTwoTransaction([]int{100, 30, 15, 10, 8, 25, 80})) // 72
	fmt.Println(MaxTwoTransaction([]int{90, 80, 70, 60, 50}))         // 0

	fmt.Println("\nMultiple Transactions")
	fmt.Println(MultipleTransaction([]int{100, 180, 260, 310, 40, 535, 695})) // 865
	fmt.Println(MultipleTransaction([]int{1, 5, 3, 8, 12}))                   // 13
	fmt.Println(MultipleTransaction([]int{5, 4, 3, 2, 1}))                    // 0

	fmt.Println("\nTransaction Fee")
	fmt.Println(TransactionWithFee([]int{6, 1, 7, 2, 8, 4}, 2)) // 8
	fmt.Println(TransactionWithFee([]int{7, 1, 5, 3, 6, 4}, 1)) // 5
	fmt.Println(TransactionWithFee([]int{1, 3, 2, 8, 4, 9}, 2)) // 8
}
