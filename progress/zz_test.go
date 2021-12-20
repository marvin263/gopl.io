package progress

import (
	"fmt"
	"testing"
)

func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= numRows {
		return s
	}

	charCountInVCol := numRows + (numRows - 2)
	vColCount := len(s) / charCountInVCol
	charCountLeft := len(s) - (vColCount * charCountInVCol)
	if charCountLeft > 0 {
		vColCount++
	}
	// 每一个虚拟列，跨越实际列的数量是 (numRows - 1)
	colCount := vColCount * (numRows - 1)

	bytes := make([][]byte, numRows, numRows)
	for idx, _ := range bytes {
		bytes[idx] = make([]byte, colCount, colCount)
	}

	idx := 0
	for vc := 0; vc < vColCount; vc++ {
		// 每一个虚拟列，跨越实际列的数量是 (numRows - 1)
		col := vc * (numRows - 1)
		// 填充竖直的
		for r := 0; r < numRows && idx < len(s); r++ {
			bytes[r][col] = s[idx]
			idx++
		}
		// 填充  左下角 延伸 到右上角
		for r, acol := numRows-1-1, 1; acol <= (numRows-2) && idx < len(s); r, acol = r-1, acol+1 {
			bytes[r][col+acol] = s[idx]
			idx++
		}
	}

	var rst []byte
	for _, aRow := range bytes {
		for _, c := range aRow {
			if c != 0 {
				rst = append(rst, c)
			}
		}
	}
	return string(rst)
}

func TestConvert(t *testing.T) {
	s1 := "PAYPALISHIRING"
	n1 := 3
	fmt.Printf("%v-->%v\n", s1, convert(s1, n1))

	s2 := "PAYPALISHIRING"
	n2 := 4
	fmt.Printf("%v-->%v\n", s2, convert(s2, n2))

	s3 := "A"
	n3 := 1
	fmt.Printf("%v-->%v\n", s3, convert(s3, n3))
}
