package basic

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Int()

	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	day := []string{"Monday", "White house", "Google"}[abs(x) % 3]

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("It's another day")
	}
}


func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
