package insertion

import (
	"testing"

	internal "github.com/lipinskipawel/orst/internal"
)

func TestInsertionSort(t *testing.T) {
	input := []int{-1, 0, 3, 1, 0, 80}
	output := []int{-1, 0, 0, 1, 3, 80}

	result := sort(input)

	if !internal.TheSame(result, output) {
		t.Errorf("got %v, want %v", result, output)
	}
}
