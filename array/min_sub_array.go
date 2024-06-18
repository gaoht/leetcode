package array

// MinSubArrayLen 209
func MinSubArrayLen2(a []int, target int) int {
	minLength := 99999999999999
	for i := 0; i < len(a); i++ {
		sum := 0
		length := 1
		for j := i; j < len(a); j++ {
			sum += a[j]
			if sum >= target {
				if length < minLength{
					minLength = length
				}
				break
			}
			length++
		}
	}
	if minLength == 99999999999999 {
		return 0
	}
	return minLength
}

func MinSubArrayLen(a []int, target int) int {
	i := 0
	sum := 0
	minLength := 99999999999999
	for	j := 0; j < len(a); j++ {	
		sum += a[j];	
		for (sum >= target) {
			length := j - i + 1
			if length < minLength {
				minLength = length
			} 
			sum -= a[i]
			i++
		}
	}
	if minLength == 99999999999999 {
		return 0
	}
	return minLength

}