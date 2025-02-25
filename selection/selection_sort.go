package selection

import "github.com/lipinskipawel/orst/internal"

// Sorts numbers based on the selection sort algorithm.
// It does not use additional memory.
// It works by finding the smallest element in the slice
// and stick it to the front of the slice, then find the
// next smallest element and stick it in second position,
// and so on.
func sort(numbers []int) []int {
	if numbers == nil {
		return nil
	}

	for i := range numbers {

		for j := i + 1; j < len(numbers); j++ {

			smallestInRest := i
			if numbers[j] > numbers[smallestInRest] {
				smallestInRest = j
			}

			if j != smallestInRest {
				internal.Swap(numbers, smallestInRest, j)
			}
		}
	}

	return numbers
}
