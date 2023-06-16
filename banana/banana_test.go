package banana_test

import (
	"testing"

	"github.com/sutin1234/go-tutorial/banana"
)

func TestShouldReturn_1_1(t *testing.T) {
	input := banana.Solution("NAABXXAN")
	if input != 1 {
		t.Error("should be except 1 actual ", input)
	}
}
func TestShouldReturn_2(t *testing.T) {
	input := banana.Solution("NAANAAXNABABYNNBZ")
	if input != 2 {
		t.Error("should be except 2 actual ", input)
	}
}
