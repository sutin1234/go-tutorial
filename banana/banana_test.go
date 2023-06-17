package banana_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutin1234/go-tutorial/banana"
)

type TestCases struct {
	caseName string
	input    string
	expected float64
}

func TestBanana(t *testing.T) {
	cases := []TestCases{
		{caseName: "NAABXXAN = 1", input: "NAABXXAN", expected: 1},
		{caseName: "NAANAAXNABABYNNBZ = 2", input: "NAANAAXNABABYNNBZ", expected: 2},
		{caseName: "QABAAWOBL = 0", input: "QABAAWOBL", expected: 0},
	}

	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			r := banana.Solution(c.input)
			if r != c.expected {
				assert.Equal(t, c.expected, c.input)
			}
		})
	}
}
