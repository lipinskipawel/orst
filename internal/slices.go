package internal

func Swap(slice []int, left int, right int) {
	tmp := slice[left]
	slice[left] = slice[right]
	slice[right] = tmp
}
