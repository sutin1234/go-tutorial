package game_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sutin1234/go-tutorial/game"
)

type TestCases struct {
	caseName  string
	input1    string
	input2    string
	expected1 string
	expected2 string
}

func TestGame(t *testing.T) {
	cases := []TestCases{
		{caseName: "20:00", input1: "20:00", input2: "22:00", expected1: "20:00", expected2: "22:00"},
	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			t1, t2 := game.Solution(c.input1, c.input2)
			assert.Equal(t, t1, c.expected1)
			assert.Equal(t, t2, c.expected2)
		})
	}

}
