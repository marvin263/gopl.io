package focus

import (
	"fmt"
	"strings"
	"testing"
)

func TestFindOcurrences(t *testing.T) {
	// ["we","we","will"]
	fmt.Println(findOcurrences("we we we we will rock you", "we", "we"))
	// []
	fmt.Println(findOcurrences("obo jvezipre obo jnvavldde jvezipre jvezipre jnvavldde jvezipre jvezipre jvezipre y jnvavldde jnvavldde obo jnvavldde jnvavldde obo jnvavldde jnvavldde jvezipre", "jnvavldde", "y"))
	// ["student"] 
	fmt.Println(findOcurrences("alice is aa good girl she is a good student", "a", "good"))
	// ["girl","student"]
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	// ["we","rock"]
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
}

func findOcurrences(text string, first string, second string) []string {
	ans := make([]string, 0, 10)

	f := first + " "
	s := second + " "
	flen := len(f)
	slen := len(s)
	from := 0
	for (from + flen + slen) < len(text) {
		// 先找到first
		idx1 := strings.Index(text[from:], f)
		if idx1 == -1 {
			break
		}
		idx1 = from + idx1
		// first前面必须得有个空格
		if idx1 > 0 && text[idx1-1] != ' ' {
			from = from + 1
			continue
		}
		// first后面跟着 second，如果不跟着，则该first是无用的
		if !strings.HasPrefix(text[idx1+flen:], s) {
			from = from + 1
			continue
		}

		// 已找到second
		idx2 := idx1 + flen

		// 顶到尾部啦
		if idx2+slen == len(text) {
			break
		} else {
			idx3 := strings.Index(text[idx2+slen:], " ")
			if idx3 == -1 {
				ans = append(ans, text[idx2+slen:])
			} else {
				idx3 = idx2 + slen + idx3
				ans = append(ans, text[idx2+slen:idx3])
			}
		}
		from = idx1 + flen
	}
	return ans
}
