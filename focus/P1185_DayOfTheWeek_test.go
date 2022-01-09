package focus

import (
	"fmt"
	"testing"
)

func TestDayOfTheWeek(t *testing.T) {
	fmt.Println(dayOfTheWeek(1, 2, 3))
}

func dayOfTheWeek(day int, month int, year int) string {
	weeks := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return weeks[0]
}
