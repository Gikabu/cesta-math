package cmath

import "math"

func lcmIterative(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcdIterative(a, b)))
}

func LCM(values ...int64) int64 {
	lcm := int64(1)
	for i, value := range values {
		if i == 1 {
			prev := values[i-1]
			lcm = lcmIterative(prev, value)
		} else if i > 1 {
			lcm = lcmIterative(lcm, value)
		} else {
			lcm = value
		}
	}
	return lcm
}
