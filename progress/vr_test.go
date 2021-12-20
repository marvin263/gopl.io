package progress

import (
	"fmt"
	"testing"
)

func equal(x, y map[string]int) bool {
	x["A1"]++
	y["B1"]++

	fmt.Println(x)
	fmt.Println(y)

	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func TestVF(t *testing.T) {
	m1 := map[string]int{"A": 0}
	m2 := map[string]int{"B": 42}
	fmt.Println(equal(m1, m2))
	fmt.Println(m1)
	fmt.Println(m2)

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	modifySlice(arr1[2:])
	modifySlice(arr2[2:])

	fmt.Println(equalArray(arr1, arr2))

	fmt.Println(arr1)
	fmt.Println(arr2)
}
func equalArray(x, y [3]int) bool {
	x[0] = 99
	y[0] = 999
	fmt.Println(x)
	fmt.Println(y)

	for idx, val := range x {
		if y[idx] != val {
			return false
		}
	}
	return true
}

func modifySlice(x []int) {
	for idx, val := range x {
		x[idx] = val + 100
	}
}
