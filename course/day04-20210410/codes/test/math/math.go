package math

func Add(l, r int) int {
	return l + r
}

func Sub(l, r int) int {
	return l - r
}

func Mult(l, r int) int {
	return l * r
}

func Fact(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return n * Fact(n-1)
}
