package quick

import (
	"github.com/lipinskipawel/orst/internal"
)

// Sorts numbers based on the quick sort algorithm.
//
// It does not use additional memory.
// It works by finding the smallest element in the slice
// and stick it to the front of the slice, then find the
// next smallest element and stick it in second position,
// and so on.
func sort(numbers []int) []int {
	if numbers == nil {
		return nil
	}

	quicksort(numbers)

	return numbers
}

func quicksort(numbers []int) {
	switch len(numbers) {
	case 0, 1:
		return
	case 2:
		if numbers[0] > numbers[1] {
			internal.Swap(numbers, 0, 1)
		}
		return
	default:
	}

	pivot := numbers[0]
	rest := numbers[1:]
	left := 0
	right := len(rest) - 1

	for left <= right {
		if rest[left] <= pivot {
			// already on the correct side
			left += 1
		} else if rest[right] > pivot {
			// deal with underflow
			if right == 0 {
				break
			}
			// already on the correct side
			right -= 1
		} else {
			// left holds a right and right holds a left, swap them
			internal.Swap(rest, left, right)
			left += 1
			if right == 0 {
				break
			}
			right -= 1
		}
	}

	// re-align left to account for the pivot at 0
	left += 1

	// place the pivot at its final location
	internal.Swap(numbers, 0, left-1)

	quicksort(rest[:left-1])
	quicksort(rest[right+1:])
}
