package advanced

// requires go1.20+ and latest GoLand 2023+

type Runner[T any] interface {
	Run() T
	CanRun(T) bool
}

/*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
 */
type MyRunner struct{}

var _ Runner[int] = (*MyRunner)(nil)

func (r MyRunner) Run() int {
	return 0
}

func (r MyRunner) CanRun(i int) bool {
	return true
}
