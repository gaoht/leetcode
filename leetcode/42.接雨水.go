package leetcode

import "fmt"
func trap(height []int) int {
    sum := 0
    for i := 0; i < len(height); i++ {
        if i == 0 || i == len(height) - 1 {
            continue
        }
        // 找左边高的
        left := height[i]
        for j := i - 1; j >= 0; j-- {
            if left < height[j] {
                left = height[j]
            }
        }
        // 找右边最高
        right := height[i]
        for j := i + 1; j < len(height); j++ {
            if right < height[j] {
                right = height[j]
            }
        }
        h := min(left, right) - height[i];
		// fmt.Println(i, left, right, h)
        if h > 0 {
            sum += h
        }
    }
    return sum
}


func trap2(height []int) int{
    leftStack := make([]int, 0, len(height))
    rightStack := make([]int, 0, len(height))
    leftStack = append(leftStack, len(height) - 1)
    rightStack = append(rightStack, 0)
    loc := make([][2]int, len(height))
    for i := range height {
        loc[i] = [2]int{-1, -1}
    }
    for i := 1; i < len(height); i++ {
        for len(rightStack) > 0 && height[rightStack[len(rightStack) - 1]] < height[i] {
            top := rightStack[len(rightStack) - 1]
            loc[top][1] = i
            rightStack = rightStack[0 : len(rightStack) - 1]
        }
        rightStack = append(rightStack, i)
    }
	fmt.Println(loc)

    for i := len(height) - 2; i >= 0; i-- {
        for len(leftStack) > 0 && height[leftStack[len(leftStack) - 1]] < height[i] {
            top := leftStack[len(leftStack) - 1]
            loc[top][0] = i
            leftStack = leftStack[0 : len(leftStack) - 1] 
        }
        leftStack = append(leftStack, i)
    }
	
    sum := 0
    for i := 1; i < len(height) - 1; i++ {
        left := loc[i][0]
		right := loc[i][1]
		if left == -1 || right == -1 {
			continue
	   }
		for loc[left][0] != -1 {
			left = loc[left][0]
		}
		for loc[right][1] != -1 {
			right = loc[right][1]
		}
        
        
        sum += min(height[left], height[right]) - height[i]
    }
    return sum
}
