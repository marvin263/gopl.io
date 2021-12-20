package main

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestName(t *testing.T) {
	m1 := make([]TreeNode, 1, 2)
	// []TreeNode
	fmt.Printf("\n%T\n", m1)
	// {0 <nil> <nil>}
	// 也就是说：非指针类型，永远没有nil
	fmt.Printf("%v\n", m1[0])

	m2 := make([](*TreeNode), 4, 6)
	// [](*TreeNode)
	fmt.Printf("\n%T\n", m2)
	// <nil>
	// 也就是说：数组的元素是 *TreeNode，而指针*TreeNode盛放的可以为nil
	// 注意：make搞出来的，m2绝不是指针哦。数组本身是 m2
	fmt.Printf("%v\n", m2[0])

	n1 := new(TreeNode)
	// *TreeNode，是的，new返回的是指针哦
	fmt.Printf("\n%T\n", n1)
	// &{0 <nil> <nil>}，对地址输出时，可能主动寻找数据本身了
	fmt.Printf("%v\n", n1)

	n2 := new([](*TreeNode))
	// *[]*main.TreeNode，数组的元素是*TreeNode
	// 注意：new搞出来的，n2就是指针
	fmt.Printf("\n%T\n", n2)
	fmt.Printf("%v\n", n2)
	fmt.Printf("%v\n", *n2)

	// 长度为5，元素是空串
	var arr [5]string
	fmt.Printf("\n%T\n", arr)
	fmt.Printf("%v\n", arr)

	// nil
	var theSlice []string
	fmt.Printf("\n%T\n", theSlice)
	fmt.Printf("%v\n", (theSlice == nil))

	// nil
	var theMap map[string]int
	fmt.Printf("\n%T\n", theMap)
	fmt.Printf("%v\n", (theMap == nil))

	n55 := new([]*TreeNode)
	n555 := new(*[]TreeNode)
	fmt.Printf("%T, %T\n", n55, n555)

}
func setValue(arr []int) {
	for idx, _ := range arr {
		arr[idx] = (idx * 10)
	}
	arr = append(arr, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9)
	fmt.Printf("inside setValue, %v\n", arr)
}

func setValueP(arr *[]int) {
	for idx, _ := range *arr {
		(*arr)[idx] = (idx * 10)
	}
	*arr = append(*arr, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9)
	fmt.Printf("inside setValueP, %v\n", arr)
}
