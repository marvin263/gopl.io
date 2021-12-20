package geometry

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	// 包级别的函数
	fmt.Println(Distance(p, q))

	// Point类下声明的 Point.Distance方法
	fmt.Println(p.Distance(q))
}
