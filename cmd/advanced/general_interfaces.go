package advanced

// requires go1.20+ and latest GoLand 2023+

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}
