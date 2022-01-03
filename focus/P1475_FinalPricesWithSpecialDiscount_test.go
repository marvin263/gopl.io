package focus

import (
	"fmt"
	palias "gopl.io/ch6/intset"
	"testing"
)

func finalPrices(prices []int) []int {
	rst := make([]int, len(prices), len(prices))
	// 单调递增
	ss := make([]int, 0, len(prices))
	s := &ss

	for i, curPrice := range prices {
		// 当前商品的价格，低
		for !Empty(s) && curPrice <= prices[Peek(s)] {
			topIdx := Pop(s)
			rst[topIdx] = prices[topIdx] - curPrice
		}
		Push(s, i)
	}
	for !Empty(s) {
		topIdx := Pop(s)
		rst[topIdx] = prices[topIdx]
	}
	return rst
}

func Empty(s *[]int) bool {
	return len(*s) == 0
}

func Push(s *[]int, v int) {
	*s = append(*s, v)
	//TODO:还需要返回这个 s 吗？
}

func Peek(s *[]int) int {
	return (*s)[len(*s)-1]
}

func Pop(s *[]int) int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func TestFinalPrices(t *testing.T) {
	var a palias.IntSet
	a.Has(1)
	a.OtherMethod(5)
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3})) //[4,2,4,2,3]
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5})) //[1,2,3,4,5]
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))   //[9,0,1,6]
}
