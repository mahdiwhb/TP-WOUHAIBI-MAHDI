package TPWOUHAIBIMAHDI

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
			return 0
	}

	min_price := prices[0]
	max_profit := 0

	for i := 1; i < len(prices); i++ {
			
			if prices[i] < min_price {
					min_price = prices[i]
			} else if prices[i] - min_price > max_profit {
					max_profit = prices[i] - min_price
			}
	}

	return max_profit
}


