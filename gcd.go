package cmath

func gcdIterative(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GCD(values ...int64) int64 {
	gcd := int64(1)
	for i, value := range values {
		if i == 1 {
			prev := values[i-1]
			gcd = gcdIterative(prev, value)
		} else if i > 1 {
			gcd = gcdIterative(gcd, value)
		} else {
			gcd = value
		}
	}
	return gcd
}
