package focus

import (
	"fmt"
	"testing"
)

func TestSlowestKey(t *testing.T) {
	fmt.Println(slowestKey([]int{9, 29, 49, 50}, "cbcd"))
	fmt.Println(slowestKey([]int{12, 23, 36, 46, 62}, "spuda"))
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxChar := uint8('a')
	prevRTime, maxDuration := 0, -1
	for idx, rTime := range releaseTimes {
		duration := rTime - prevRTime
		if duration > maxDuration {
			maxDuration = duration
			maxChar = keysPressed[idx]
		} else if duration < maxDuration {
			// do noting
		} else {
			if keysPressed[idx] > maxChar {
				maxChar = keysPressed[idx]
			}
		}
		prevRTime = rTime
	}
	return maxChar
}
