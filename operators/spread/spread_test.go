package operattors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	operattors "github.com/sutin1234/go-tutorial/operators/spread"
)

type TestCases struct {
	caseName string
	input   []int
	expected int
}

func TestSumSpread(t *testing.T){
	cases := []TestCases{
		{caseName: "case_1", input: []int{1,2,3,4,5}, expected: 15},
		{caseName: "case_2", input: []int{1,2,9,4,5}, expected: 21},
		{caseName: "case_3", input: []int{1,2,3,2,5}, expected: 13},
	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			v := operattors.Sum(c.input...)
			if v != c.expected {
				assert.Equal(t, c.expected, c.input)
			}
		})
	}
}