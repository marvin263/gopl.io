package progress

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestKindsOfString(t *testing.T) {
	pb, err := strconv.ParseBool("True")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pb)
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strings.Repeat("Hello", 3))
}
