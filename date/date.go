package date

import "time"

var now = time.Now()

func Now() time.Time {
	return now
}
func CurrentDate() string {
	return now.Format("2006-01-02")
}
func DateFormat(date string, format string) string {
	dt := date
	time, err := time.Parse(format, dt)
	if err != nil {
		panic(err)
	}
	return time.Format(format)
}
