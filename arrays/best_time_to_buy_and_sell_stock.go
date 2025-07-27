package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	maxProfit := 0
	for _, price := range prices {
		currProfit := price - min
		if currProfit > maxProfit {
			maxProfit = currProfit
		}
		if price < min {
			min = price
		}
	}
	return maxProfit
}
