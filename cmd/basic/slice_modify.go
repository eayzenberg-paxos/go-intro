package basic

import "fmt"

func main() {
	// Example 1: Modifying a slice element
	nums := []int{1, 2, 3, 4, 5}
	nums[2] = 10
	fmt.Println(nums) // Output: [1 2 10 4 5]

	// Example 2: Appending elements to a slice
	letters := []string{"a", "b", "c"}
	letters = append(letters, "d")
	fmt.Println(letters) // Output: [a b c d]

	// Example 3: Slicing a slice
	nums = []int{1, 2, 3, 4, 5}
	newNums := nums[1:4]
	newNums[0] = 10
	fmt.Println(nums)    // Output: [1 10 3 4 5]
	fmt.Println(newNums) // Output: [10 3 4]

}
