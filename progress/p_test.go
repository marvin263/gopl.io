package progress

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var a int = 1
	var b *int = &a            // 0xc00000a320，b作为一个指针，b指针本身的值是 变量a的地址，b指针指向了 变量a的值
	var c **int = &b           // 0xc000006038，c作为一个指针，c指针本身的值是 “b指针本身”的地址，c指针指向了 “b指针本身”的地址
	var x int = *b             // b必须是个指针，然后，这根指针所指向的内容，也就是 变量a的值
	fmt.Println("a = ", a)     // 1
	fmt.Println("&a = ", &a)   // 0xc00000a320
	fmt.Println("*&a = ", *&a) // 变量a的地址 所指向的内容，即变量a的值，1
	fmt.Println("b = ", b)     // b这根指针 本身，0xc00000a320
	fmt.Println("&b = ", &b)   //  b这根指针 的地址，0xc000006038
	fmt.Println("*&b = ", *&b) // b这根指针 的地址 所指向的内容，也就是 b指针本身，即：0xc00000a320
	fmt.Println("*b = ", *b)   // b是个指针，然后，这根指针所指向的内容，也就是 变量a 的值 1
	fmt.Println("c = ", c)     // c是个指针，然后，这根指针本身的内容（不是其指向的内容），也就是b的地址 0xc000006038
	fmt.Println("*c = ", *c)   // c必须是个指针，然后，这根指针所指向的内容，也就是 b的地址 所指向的内容，即：变量a 的地址 0xc00000a320
	fmt.Println("&c = ", &c)   //c是个指针，这根指针的地址 0xc000006040
	fmt.Println("*&c = ", *&c) //c是个指针，这根指针的地址 所指向的内容，还是，c指针本身，0xc000006038
	fmt.Println("**c = ", **c) //
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x) // 1
}
