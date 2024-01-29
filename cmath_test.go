package cmath

import "testing"

func TestFractions_MakeSumWhole(t *testing.T) {
	fs := Fractions{
		FractionFromFloatDefault(0.65),
		FractionFromFloatDefault(0.23),
		FractionFromFloatDefault(0.06),
		FractionFromFloatDefault(0.06),
	}

	t.Logf("Before MakeSumWhole, Sum: %s", fs.Sum().String())
	fs.MakeSumWhole()
	t.Logf("After MakeSumWhole, Sum: %s", fs.Sum().String())
	if !fs.IsSumWhole() {
		t.Errorf("IsSumWhole: expected true got false")
	}
}
