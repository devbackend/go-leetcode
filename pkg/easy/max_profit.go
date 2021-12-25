package easy

// MaxProfit for https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {
	var max int

	minPrice := prices[0]

	for k := 0; k < len(prices); k++ {
		profit := prices[k] - minPrice
		if profit > max {
			max = profit
		}

		if prices[k] < minPrice {
			minPrice = prices[k]
		}
	}

	return max
}
