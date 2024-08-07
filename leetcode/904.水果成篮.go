package leetcode
//  [904] 水果成篮

func totalFruit(fruits []int) int {
	min := 0
	baskets := [2]int{}
	baskets_len := 0
	count := 0

	for i := 0; i < len(fruits); i++{
		f := fruits[i]
		if baskets_len == 0 {
			baskets[0] = f
			baskets_len = 1
			count++
		} else if baskets_len == 1 {
			if baskets[0] != f {
				baskets[1] = f
				baskets_len = 2
			}
			count++
		} else if baskets_len == 2 {
			if baskets[0] == f || baskets[1] == f{
				count++
			} else {
				baskets_len = 2
				baskets[0] = fruits[i - 1]
				baskets[1] = f
				count = 2
				j := i - 2
				for j >= 0 && fruits[j] == fruits[i - 1] {
					count++
					j--
				}
			}
		}
		if count > min {
			min = count
		}
	}
	return min
}

