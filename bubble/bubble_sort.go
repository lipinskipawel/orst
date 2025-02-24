package bubble

import (
	internal "github.com/lipinskipawel/orst/internal"
)

// Sorts numbers based on the bubble sort algorithm.
// It works by checking each two numbers and swaps them if necessary.
// Continues until there are no more swaps to be made.
func sort(numbers []int) []int {
	if numbers == nil {
		return nil
	}

	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(numbers); i++ {
			if numbers[i-1] > numbers[i] {
				internal.Swap(numbers, i, i-1)
				swapped = true
			}
		}
	}
	// for i := 0; i < len(numbers); i++ {
	// 	for j := 0; j < len(numbers); j++ {
	// 		if numbers[i] < numbers[j] {
	//			internal.Swap(numbers, i, j)
	// 		}
	// 	}
	// }
	return numbers
}
