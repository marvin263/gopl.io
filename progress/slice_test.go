package progress

import (
	"fmt"
	"testing"
)

func TestMultiDimensionSlice(t *testing.T) {
	numRows := 3
	colCount := 10
	bytes := make([][]byte, numRows, numRows)
	for idx, _ := range bytes {
		bytes[idx] = make([]byte, colCount, colCount)
	}

	Print2DimensionSlice("001-", bytes)
	for idx, aRow := range bytes {
		bytes[idx] = append(aRow, 9, 10, 11)
	}
	Print2DimensionSlice("002-", bytes)
}

func rangeVar(bslc []byte) {
	for idx, val := range bslc {
		val = 2
		fmt.Printf("%v, %v", idx, val)
	}
}
func Print2DimensionSlice(tag string, twoDimensionSlice [][]byte) {
	fmt.Printf("\"%v\"\n", tag)
	for idx, aRow := range twoDimensionSlice {
		fmt.Printf("row - %d\n", idx)
		for i, v := range aRow {
			fmt.Printf("cell[%d]=%v", i, v)
			if i < len(aRow)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n\n")
	}
}

func TestSliceRange(t *testing.T) {
	slc := make([]string, 1, 2)
	slc[0] = "你好"
	PrintSlice("01-", slc)

	slc[1] = "世界" //panic，只能设置到 len-1；cap决定是否新创建数组
	PrintSlice("02-", slc)
}
func TestAbc(t *testing.T) {
	slc := make([]string, 1, 2)
	slc[0] = "诸葛亮"
	PrintSlice("slc-1", slc)

	newslc1 := append(slc, "a")
	newslc1[0] = "诸葛村夫"
	PrintSlice("newslc-1", newslc1)
	PrintSlice("slc-1", slc)

	fmt.Println("slice扩容后")
	fmt.Println()

	newslc2 := append(slc, "x", "y")
	newslc2[0] = "卧龙先生"
	PrintSlice("newslc-2", newslc2)
	PrintSlice("slc-1", slc)
}
func PrintSlice(tag string, slc []string) {
	fmt.Printf("\"%v\"\n", tag)
	fmt.Printf("len=%d, cap=%d, pointer=%p\n", len(slc), cap(slc), slc)
	for idx, val := range slc {
		fmt.Printf("idx=%v, val=%v\n", idx, val)
	}
	fmt.Println()
}
