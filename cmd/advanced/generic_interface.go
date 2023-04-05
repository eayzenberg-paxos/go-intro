package advanced

// requires go1.20+ and latest GoLand 2023+

type Runner[T any] interface {
	Run() T
}
