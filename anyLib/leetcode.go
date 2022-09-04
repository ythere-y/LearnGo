package anyLib

func finalPrices(prices []int) []int {
	var idxes []int
	for i, val := range prices {
		if len(idxes) == 0 || prices[idxes[len(idxes)-1]] < val {
			idxes = append(idxes, i)
			continue
		}
		for len(idxes) != 0 && prices[idxes[len(idxes)-1]] >= val {
			prices[idxes[len(idxes)-1]] -= val
			idxes = idxes[:len(idxes)-1]
		}
		idxes = append(idxes, i)

	}
	return prices
}
