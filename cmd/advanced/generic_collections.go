package advanced

import "fmt"

// requires go1.20+ and latest GoLand 2023+

// Map - given a slice of type T, executes the passed in mapper function for each element in the slice, returning a slice of type R.
// The function is passed the current element, the current index and the slice itself as function arguments.
// source: https://github.com/Goldziher/go-utils/blob/main/sliceutils/sliceutils.go
func Map[T any, R any](slice []T, mapper func(value T, index int, slice []T) R) (mapped []R) {
	for i, el := range slice {
		mapped = append(mapped, mapper(el, i, slice))
	}
	return mapped
}

func main() {
	println(
		Map[int, string]([]int{1, 2, 3}, func(value int, index int, slice []int) string {
			return fmt.Sprint(value)
		}),
	)
}

/* Invalid yet

type HighOrderList[T any] []T

func (list HighOrderList[T]) mapList[V any](mapper func(T) V) HighOrderList[V] {
	out := make([]V, len(list))
	for i, t2 := range list {
		out[i] = mapper(t2)
	}
	return out
}

*/
