package greedy

import "sort"

// MinCoin
func MinCoin(coins []int, amount int) int {
	if(len(coins) == 0) {
		return -1
	}
	sort.Slice(coins, func(i, j int) bool {
		return i > j
	})
	for i := 0; i < len(coins); i++ {
		
	}
}	