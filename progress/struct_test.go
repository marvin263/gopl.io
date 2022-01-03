package progress

import (
	"fmt"
	"testing"
)

type ComplexStruct struct {
	str  string
	num  int32
	arr  [2]string
	slc  []string
	slcC []string
	m    map[string]int32
}

func TestStruct(t *testing.T) {
	cs := ComplexStruct{
		str:  "Hello",
		num:  911,
		arr:  [2]string{"ChinaSystem", "ZTE"},
		slc:  []string{"January", "February"},
		slcC: []string{"Monday", "Tuesday"},
		m:    map[string]int32{"pow": 2, "sin": 1},
	}
	fmt.Printf("1--\ncs=%p\n%v\ncs.slcC=%p, &cs.slcC=%p\n\n", &cs, cs, cs.slcC, &cs.slcC)
	changeCS(cs)
	//入参非指针时，slc, map都会改变，其他都不会
	fmt.Printf("4--\ncs=%p\n%v\ncs.slcC=%p, &cs.slcC=%p\n\n", &cs, cs, cs.slcC, &cs.slcC)
}

func changeCS(cs ComplexStruct) {
	fmt.Printf("2--\ncs=%p\n%v\ncs.slcC=%p, &cs.slcC=%p\n\n", &cs, cs, cs.slcC, &cs.slcC)
	cs.str = "蓝色星球"
	cs.num = 9527
	cs.arr[0] = "比亚迪"
	cs.slc[0] = "十二月过年了"
	cs.slcC[0] = "星期一"
	cs.slcC = append(cs.slcC, "星期三")
	cs.slcC = append(cs.slcC, "星期四")
	cs.slcC = append(cs.slcC, "星期五")
	cs.slcC = append(cs.slcC, "星期六")
	cs.m["a平方根"] = 38
	fmt.Printf("3--\ncs=%p\n%v\ncs.slcC=%p, &cs.slcC=%p\n\n", &cs, cs, cs.slcC, &cs.slcC)
}

func TestStructPointer(t *testing.T) {
	cs := ComplexStruct{
		str: "Hello",
		num: 911,
		arr: [2]string{"ChinaSystem", "ZTE"},
		slc: []string{"January", "February"},
		m:   map[string]int32{"pow": 2, "sin": 1},
	}
	fmt.Printf("1--\n%v\n", cs)
	changeCSPointer(&cs)
	fmt.Printf("2--\n%v\n", cs)

}

func changeCSPointer(cs *ComplexStruct) {
	cs.str = "蓝色星球"
	cs.num = 9527
	cs.arr[0] = "比亚迪"
	cs.slc[0] = "十二月过年了"
	cs.m["a平方根"] = 38
}
