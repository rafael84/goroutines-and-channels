package random

import "time"

func SleepMs(min int, max int) {
	time.Sleep(time.Millisecond * time.Duration(IntBetween(100, 200)))
}
