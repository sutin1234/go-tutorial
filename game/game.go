package game

import (
	"time"
)

func Solution(a string, b string) (string, string) {
	// a = start game time, b = end game time

	// extrach hour, minute
	layout := "15:04" // 24-hour HH:MM
	t1, _ := time.Parse(layout, a)
	t2, _ := time.Parse(layout, b)
	// fmt.Printf("hh: %v mm: %v\n", t1.Format(layout), t2.Format(layout))

	// cal for crrentDay
	return t1.Format(layout), t2.Format(layout)

}
