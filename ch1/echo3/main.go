// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
)

type Student struct {
	Id      uint32
	Name    string
	Address string
}

func main() {
	var str = "Hello世界"
	var count = 99
	fmt.Printf("%v, %v\n", str, count)

	var f1 = 123.789
	var f2 = 123456789.123456789
	fmt.Printf("%v, %v\n", f1, f2)

	var stu = Student{
		Id:      10036,
		Name:    "Marvin",
		Address: "中国江苏Nanjing",
	}
	fmt.Printf("%v\n", stu)
	fmt.Printf("%+v\n", stu)

	var str2 = `	仰知天文,俯察地理,中晓人和,
	明阴阳,懂八卦,晓奇门,知遁甲, 运筹帷幄之中,决胜千里之外,
	自比管仲乐毅之贤,抱膝危坐,笑傲风月,未出茅庐,先定三分天下。 `
	fmt.Printf("%#v\n", str2)
	fmt.Printf("%q\n", str2)

	fmt.Printf("%d%%的财富掌握在%d%%的人手中\n", 80, 20)
	fmt.Printf("%T, %T, %T\n", count, f1, stu)
	fmt.Printf("%x, %X\n% x, % X\n", str, str, str, str)
	fmt.Printf("%b, %c, %d\n%o, %q, %U\n", 198, 198, 198, 198, 198, 198)
	fmt.Printf("%+f, %+q", 3.1415926, "Hello	世界")

}

//!-
