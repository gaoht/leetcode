package leetcode

import (
	"fmt"
	"strings"
)
type IntensitySegments struct{
	segments []int
	weights []int
}

// NewIntensitySegments new
func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{}
}

// Add add
func (is *IntensitySegments) Add(from, to, amount int){
	// left := binarySearch(is.segments, from)
	seg := make([]int, 0)
	weights := make([]int, 0)
	if !check(from, to) {
		return
	}
	
	for i := 0; i < len(is.segments); i ++ {
		if is.segments[i] < from {
			seg, weights = appendSegWeight(seg, weights, is.segments[i], is.weights[i])
		}
		if is.segments[i] == from {
			seg, weights = appendSegWeight(seg, weights, is.segments[i], is.weights[i] + amount)
		}
		if is.segments[i] > from && is.segments[i] < to {
			if len(seg) == 0 {
				seg, weights = appendSegWeight(seg, weights, from, amount)
			} else if seg[len(seg) - 1] < from {
				seg, weights = appendSegWeight(seg, weights, from, is.weights[i - 1] + amount)
			}
			seg, weights = appendSegWeight(seg, weights, is.segments[i], is.weights[i] + amount)
		}
		// == to + 0
		// if is.segments[i] == to 	
		if is.segments[i] > to {
			if seg[len(seg) - 1] < to {
				seg, weights = appendSegWeight(seg, weights, to, is.weights[i])
			}
			seg, weights = appendSegWeight(seg, weights, is.segments[i], is.weights[i])
		}
	}
	// loop skipped from is largest
	if len(seg) == 0 || seg[len(seg) - 1] < from {
		seg, weights = appendSegWeight(seg, weights, from, amount)
		seg, weights = appendSegWeight(seg, weights, to, 0)
	} 
	if seg[len(seg) - 1] < to {
		seg, weights = appendSegWeight(seg, weights, to, 0) 
	}
	is.segments = seg
	is.weights = weights
	is.remove2Zeros()
}
func (is *IntensitySegments) remove2Zeros() {
	index := 0;
	for cursor := 0; cursor < len(is.segments); cursor++ {
		if is.weights[cursor] == 0 {
			// starts with 0
			if index == 0 {
				continue
			} else if is.weights[index - 1] == 0{
				// 2 zeros
				continue
			}
		} 
		is.segments[index] = is.segments[cursor]
		is.weights[index] = is.weights[cursor]
		index++
	}
	is.segments = is.segments[ : index]
	is.weights = is.weights[ : index]
}

func appendSegWeight(segments, amounts []int, segment, amount int) ([]int, []int){
	segments = append(segments, segment)
	amounts = append(amounts, amount)
	return segments, amounts
}

func (is *IntensitySegments) ToString() string {
	var x strings.Builder
	x.WriteString("[")
	for i, seg := range is.segments{
		x.WriteString(fmt.Sprintf("[%d,%d]", seg, is.weights[i]))
		if i != len(is.segments) - 1 {
			x.WriteString(",")
		}
	}
	x.WriteString("]")
	return x.String()
}



func check(from, to int) bool {
	// [from, to) 
	if from >= to {
		return false
	}
	return true
}


