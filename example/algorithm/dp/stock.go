package dp

// MaxProfitOnce 买卖股票最佳时机 只买卖1次
func MaxProfitOnce(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	small, m := prices[0], 0
	for i := range prices {
		if prices[i] < small {
			small = prices[i]
		} else {
			if prices[i]-small > m {
				m = prices[i] - small
			}
		}
	}
	return m
}

// MaxProfit 买卖股票最佳时机 每次只能买一次或卖一次
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	m, l := 0, len(prices)
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			m = prices[i] - prices[i-1] + m
		}
	}
	return m
}
