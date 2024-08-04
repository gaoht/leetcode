/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根 
 */
package leetcode
// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1{
		return x
	}
	left := 0
	right := x
	for left + 1 != right {
		mid := left + (right - left) / 2
		if mid * mid < x {
			left = mid
		} else if mid * mid > x {
			right = mid
		} else {
			return mid
		}
	}
	return left
}
// @lc code=end

