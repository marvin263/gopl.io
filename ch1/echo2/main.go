// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Printf("%T, %T\n", [...]string{"Hello世界"}, []string{"Hello世界"})
	fmt.Printf("%v, %v", cap([...]string{"Hello世界"}), cap([]string{"Hello世界"}))
	var values []int
	fmt.Println(len(values))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-
