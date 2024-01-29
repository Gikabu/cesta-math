package cmath

import "testing"

func TestFractions_MakeSumWhole(t *testing.T) {
	fs := Fractions{&Fraction{
		Parts: 3,
		Base:  10,
	}, &Fraction{
		Parts: 3,
		Base:  10,
	}, &Fraction{
		Parts: 31,
		Base:  100,
	}}

	t.Logf("Before MakeSumWhole, Sum: %s", fs.Sum().String())
	fs.MakeSumWhole()
	t.Logf("After MakeSumWhole, Sum: %s", fs.Sum().String())
	if !fs.IsSumWhole() {
		t.Errorf("IsSumWhole: expected true got false")
	}
}
