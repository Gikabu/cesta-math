package cmath

import "fmt"

type Fraction struct {
	Parts int64 `json:"parts"`
	Base  int64 `json:"base"`
}

type Fractions []*Fraction

const (
	ten = 1e1
)

const (
	DefaultFractionDecimals = 4
)

func getBase(decimals int64) int64 {
	return ten ^ decimals
}

func NewFraction(parts int64, per int64) *Fraction {
	return &Fraction{
		Parts: parts,
		Base:  per,
	}
}

func FractionFromFloatDefault(value float64) *Fraction {
	base := getBase(DefaultFractionDecimals)
	parts := int64(value * float64(base))
	return NewFraction(parts, base)
}

func FractionFromFloat(value float64, decimals int64) *Fraction {
	base := getBase(decimals)
	parts := int64(value * float64(base))
	return NewFraction(parts, base)
}

func (f *Fraction) Decimal() float64 {
	return float64(f.Parts) / float64(f.Base)
}

func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Parts, f.Base)
}

func (fs *Fractions) basesLCM() int64 {
	var bases []int64
	for _, fraction := range *fs {
		bases = append(bases, fraction.Base)
	}
	return LCM(bases...)
}

func (fs *Fractions) MakeSumWhole() {
	sum := fs.Sum()
	if sum.Base == sum.Parts || len(*fs) < 1 {
		return
	}

	factor := sum.Base / (*fs)[0].Base
	diff := sum.Base - sum.Parts
	(*fs)[0].Base *= factor
	(*fs)[0].Parts *= factor
	(*fs)[0].Parts += diff
}

func (fs *Fractions) IsSumWhole() bool {
	sum := fs.Sum()
	return sum.Base == sum.Parts
}

func (fs *Fractions) Sum() *Fraction {
	sumParts := int64(0)
	basesLCM := fs.basesLCM()
	for _, fraction := range *fs {
		sumParts += fraction.Parts * (basesLCM / fraction.Base)
	}
	return NewFraction(sumParts, basesLCM)
}
