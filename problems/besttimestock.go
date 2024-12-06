package problems

// maxProfit calculates the maximum profit from buying and selling a stock
func maxProfit(prices []int) int {
	// Handle empty or single-element array
	if len(prices) < 2 {
		return 0
	}

	// Track the lowest price seen so far
	minPrice := prices[0]

	// Track maximum potential profit
	maxProfit := 0

	// Iterate through prices starting from second element
	for i := 1; i < len(prices); i++ {
		// Update lowest price if current price is lower
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Calculate potential profit at current price
			currentProfit := prices[i] - minPrice

			// Update maximum profit if current profit is higher
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	// Return maximum possible profit
	return maxProfit
}
