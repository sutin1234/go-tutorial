package hello_test

import (
	"testing"

	"github.com/sutin1234/go-tutorial/hello"
)

func TestShouldVer_1_20_byDefault(t *testing.T) {
	p := hello.GoParams{Name: "", Version: ""}
	v := hello.GoVersion(p)
	if v != "golang 1.20" {
		t.Error("should be except golang 1.20 actual ", p.Name+p.Version)
	}
}
func TestShouldVer_1_20_byInput(t *testing.T) {
	p := hello.GoParams{Name: "golang", Version: "1.20"}
	v := hello.GoVersion(p)
	if v != p.Name+" "+p.Version {
		t.Error("should be except golang 1.20 actual ", p.Name+p.Version)
	}
}
