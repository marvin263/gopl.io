package progress

import (
	"fmt"
	"testing"
)

type NewYear struct {
	score     float64
	golang    string
	wishCount uint64
}

func TestArray(t *testing.T) {
	arr := [5]string{"a0", "a1", "a2", "a3", "a4"}
	fmt.Printf("arr=%v, &arr=%p\n", arr, &arr)
}

func TestArrayAndSlice(t *testing.T) {
	arr := [5]string{"a0", "a1", "a2", "a3", "a4"}
	for idx, _ := range arr {
		fmt.Printf("arr[%d]=%v, &arr[%d]=%p\n", idx, arr[idx], idx, &(arr[idx]))
	}
	fmt.Println()

	slc0 := arr[0:]
	fmt.Printf("slc0=%p, &slc0=%p\n", slc0, &slc0)

	slc2 := arr[2:]
	fmt.Printf("slc2=%p, &slc2=%p\n", slc2, &slc2)
}

func TestFormalParam(t *testing.T) {
	slc := []string{"a0", "a1", "a2", "a3", "a4"}
	fmt.Printf("slc=%p, &slc=%p, slc=%v\n", slc, &slc, slc)
	setValue(slc)
	fmt.Printf("slc=%p, &slc=%p, slc=%v\n", slc, &slc, slc)
}

func setValue(slc2 []string) {
	fmt.Printf("slc2=%p, &slc2=%p, slc2=%v\n", slc2, &slc2, slc2)
	slc2[0] = "诸葛村夫"
	slc2[1] = "织席贩履"
	slc2[2] = "大耳刘贼"
}
