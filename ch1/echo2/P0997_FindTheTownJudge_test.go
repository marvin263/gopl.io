package main

import (
	"fmt"
	"testing"
)

func findJudge(n int, trust [][]int) int {
	trustedCountArr := make([]int, n+1, n+1)
	alreadyTrustSomeOneArr := make([]bool, n+1, n+1)

	for _, whoTrustWho := range trust {
		someOne := whoTrustWho[0]
		beTrustedOne := whoTrustWho[1]
		alreadyTrustSomeOneArr[someOne] = true
		if someOne != beTrustedOne {
			trustedCountArr[beTrustedOne]++
		}
	}

	c, judge := 0, 0
	for someOne := 1; someOne < len(trustedCountArr); someOne++ {
		trustedCount := trustedCountArr[someOne]
		if trustedCount == n-1 && !alreadyTrustSomeOneArr[someOne] {
			judge = someOne
			c++
		}
	}

	if c == 1 {
		return judge
	} else {
		return -1
	}
}

func TestFindJudge(t *testing.T) {

	fmt.Println(findJudge(2, [][]int{{1, 2}})) //2

	fmt.Println(findJudge(3, [][]int{
		{1, 3},
		{2, 3}})) //3

	fmt.Println(findJudge(3, [][]int{
		{1, 3},
		{2, 3},
		{3, 1}})) //-1

	fmt.Println(findJudge(3, [][]int{
		{1, 2},
		{2, 3}})) //-1

	fmt.Println(findJudge(4, [][]int{
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{4, 3}})) //3
}
