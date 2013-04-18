package graph

func Min(x, y int) int {
	if x < y {
		return x
	} else if y < x {
		return y
	}
	return x
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
