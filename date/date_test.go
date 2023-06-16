package date_test

import (
	"testing"

	"github.com/sutin1234/go-tutorial/date"
)

func TestShouldBeCurrentDate(t *testing.T) {
	d := date.Now().Format("2006-01-02")
	n := date.CurrentDate()
	if d != n {
		t.Error("should be except "+n+" actual ", d)
	}
}
func TestShouldBeNotCurrentDate(t *testing.T) {
	d := date.Now().Format("2006/01/02")
	n := date.CurrentDate()
	if d == n {
		t.Error("should be except "+n+" actual ", d)
	}
}
