// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Boiling prints the boiling point of water.
package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

const boilingF = 212.0

func main() {
	values := []int{1, math.MaxInt32, 2, math.MaxInt32, 3}
	root := buildTree(values)
	balanced := isBalanced(root)
	fmt.Println(balanced)
	abc()
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
	fmt.Println(truncateSentence("chopper is not a tanuki", 5))
	// Output:
	// boiling point = 212°F or 100°C
}

func truncateSentence(s string, k int) string {
	count := 0
	finalIndex := 0
	for i, c := range s {
		finalIndex = i
		if c == ' ' {
			count++
		}
		if count == k {
			break
		}
	}
	if count != k {
		finalIndex++
	}
	return s[:finalIndex]
}

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == ')' || c == ']' || c == '}' {
			if len(stack) != 0 {
				c2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if match(c2, c) {
					continue
				} else {
					return false
				}
			} else {
				return false
			}

		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

func match(c1, c2 rune) bool {
	return (c1 == ')' && c2 == '(') ||
		(c1 == '[' && c2 == ']') ||
		(c1 == '{' && c2 == '}') ||
		(c2 == ')' && c1 == '(') ||
		(c2 == '[' && c1 == ']') ||
		(c2 == '{' && c1 == '}')
}

func detectCapitalUse(word string) bool {
	var firstLetterCapital = false
	var uppercaseCount = 0

	for idx, aChar := range word {
		if aChar >= 'A' && aChar <= 'Z' {
			uppercaseCount++
			if idx == 0 {
				firstLetterCapital = true
			}
		}
	}

	//全小写
	return (uppercaseCount == 0) ||
		//全大写
		uppercaseCount == utf8.RuneCountInString(word) ||
		// 首字母大写
		(firstLetterCapital && uppercaseCount == 1)
}

func abc() {
	var str = "H世界"
	for idx, _ := range str {
		fmt.Printf("idx=%d, aChar=%T\n", idx, str[idx])
	}
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%s, %d, %d\n", str[i:i+size], size, r)
		i += size
	}
	stringSlice := []string{"PowerPoint", "Excel", "Word"}
	for idx, str := range stringSlice {
		fmt.Printf("%T, idx=%T, %d; str=%T, %s\n", stringSlice, idx, idx, str, str)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 如何传递slice指针
func buildTree(values []int) *TreeNode {
	fmt.Println(values)
	if len(values) == 0 {
		return nil
	}
	half := len(values) / 2
	parents := make([]*TreeNode, half, half)
	root := TreeNode{
		Val:   values[0],
		Left:  nil,
		Right: nil,
	}
	parents[0] = &root

	for i := 1; i <= half; i++ {
		p := parents[i-1]
		if p == nil {
			continue
		}

		leftIdx := 2*i - 1
		rightIdx := 2 * i
		if values[leftIdx] != math.MaxInt32 {
			left := TreeNode{
				Val:   values[leftIdx],
				Left:  nil,
				Right: nil,
			}
			p.Left = &left
			if leftIdx <= half {
				parents[leftIdx] = &left
			}
		}

		if rightIdx <= len(values)-1 && values[rightIdx] != math.MaxInt32 {
			right := TreeNode{
				Val:   values[rightIdx],
				Left:  nil,
				Right: nil,
			}
			p.Right = &right
			if rightIdx <= half {
				parents[leftIdx] = &right
			}
		}
	}

	return &root
}

func isBalanced(root *TreeNode) bool {
	//TODO: map的key用什么
	//TODO:怎么创建，new，还是make
	//TODO:slice取时，取地址还是可以取到原值
	//TODO:struct能当map的key使用
	//TODO:struct能当map的key使用
	var mapHeight = &map[*TreeNode]int{}
	return checkBalanced(root, mapHeight)
}

func checkBalanced(root *TreeNode, mapHeight *map[*TreeNode]int) bool {
	if root == nil {
		return true
	}
	leftHeight := heightOfNode(root.Left, mapHeight)
	rightHeight := heightOfNode(root.Right, mapHeight)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}
	return checkBalanced(root.Left, mapHeight) && checkBalanced(root.Right, mapHeight)
}

func heightOfNode(n *TreeNode, mapHeight *map[*TreeNode]int) int {
	if n == nil {
		return 0
	}
	if (*mapHeight)[n] > 0 {
		return (*mapHeight)[n]
	}
	var leftHeight, rightHeight int
	if n.Left == nil {
		leftHeight = 1
	} else {
		leftHeight = 1 + heightOfNode(n.Left, mapHeight)
	}

	if n.Right == nil {
		rightHeight = 1
	} else {
		rightHeight = 1 + heightOfNode(n.Right, mapHeight)
	}
	var heightOfN int
	if leftHeight > rightHeight {
		heightOfN = leftHeight
	} else {
		heightOfN = rightHeight
	}
	(*mapHeight)[n] = heightOfN
	return heightOfN
}

//!-
