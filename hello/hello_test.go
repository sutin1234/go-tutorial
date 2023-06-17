package hello_test

import (
	"testing"

	"github.com/sutin1234/go-tutorial/hello"
)

func TestVersion(t *testing.T) {

	type testCases struct {
		caseName string
		params   hello.GoParams
		expected string
	}
	cases := []testCases{
		{caseName: "by_default", params: hello.GoParams{Name: "", Version: ""}, expected: "golang 1.20"},
		{caseName: "by_input", params: hello.GoParams{Name: "golang", Version: "1.20"}, expected: "golang 1.20"},
	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			p := c.params
			v := hello.GoVersion(p)
			if v != c.expected {
				t.Errorf("should be except %v %v actual %v %v", p.Name, p.Version, p.Name, p.Version)
			}
		})
	}
}
