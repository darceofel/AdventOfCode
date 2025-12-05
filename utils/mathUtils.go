package utils

func PowInt(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		a *= a
		b >>= 1
	}
	return p
}
