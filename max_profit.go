package main

func maxProfit(prices []int) int {
	l := len(prices)
	switch l {
	case 0, 1:
		return 0
	}

	min, max, maxProfit := prices[0], prices[0], 0

	for _, price := range prices[1:] {
		if price < min {
			min = price
			continue
		}

		if price > max {
			max = price
		}

		profit := price - min
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
