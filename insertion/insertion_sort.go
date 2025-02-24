package insertion

import "github.com/lipinskipawel/orst/internal"

// Sorts numbers based on the insertion sort algorithm.
// It works by divding slice into sorted and not sorted parts
// [ sorted | not sorted ], example [ 1 3 4 | 2 ]
// Take first element from not sorted and walks backwards to place
// it in correct order.
func sort(numbers []int) []int {
	if numbers == nil {
		return nil
	}

	unsorted := 1
	for unsorted < len(numbers) {

		i := unsorted
		for i > 0 && numbers[i-1] > numbers[i] {
			internal.Swap(numbers, i, i-1)
			i -= 1
		}
		unsorted++
	}
	return numbers
}
