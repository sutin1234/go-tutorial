package hello_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutin1234/go-tutorial/hello"
	Struct "github.com/sutin1234/go-tutorial/struct"
)

type TestCases struct {
	caseName string
	params   Struct.GoParams
	expected string
}

func TestVersion(t *testing.T) {
	cases := []TestCases{
		{caseName: "by_default", params: Struct.GoParams{Name: "", Version: ""}, expected: "golang 1.20"},
		{caseName: "by_input", params: Struct.GoParams{Name: "golang", Version: "1.20"}, expected: "golang 1.20"},
	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			v := hello.GoVersion(c.params)
			if v != c.expected {
				assert.Equal(t, c.expected, c.params)
			}
		})
	}
}
