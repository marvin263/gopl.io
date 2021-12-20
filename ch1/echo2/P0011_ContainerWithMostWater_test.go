package main

import (
	"fmt"
	"math"
	"testing"
)

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	leftPointer, rightPointer := 0, len(height)-1

	tmpMaxArea, maxAreaLeft, maxAreaRight := 0, -1, -1

	for leftPointer < rightPointer {
		leftV, rightV := height[leftPointer], height[rightPointer]
		area := (rightPointer - leftPointer) * int(math.Min(float64(leftV), float64(rightV)))
		if tmpMaxArea < area {
			tmpMaxArea, maxAreaLeft, maxAreaRight = area, leftPointer, rightPointer
		}

		if leftV <= rightV {
			leftPointer++
		} else {
			rightPointer--
		}
	}
	_, _ = maxAreaLeft, maxAreaRight

	return tmpMaxArea
}

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) //49
	fmt.Println(maxArea([]int{1, 1}))                      //1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             //16
	fmt.Println(maxArea([]int{1, 2, 1}))                   //2
}
