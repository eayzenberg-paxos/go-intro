package advanced

// requires go1.20+ and latest GoLand 2023+

func AnyMax[T numbers](t1 T, t2 T) T {
	if t1 >= t2 {
		return t1
	}
	return t2
}
