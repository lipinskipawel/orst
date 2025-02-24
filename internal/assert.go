package internal

func TheSame(a []int, b []int) bool {
	if a != nil && b != nil {
		if len(a) != len(b) {
			return false
		}

		for i, n := range b {
			if a[i] != n {
				return false
			}
		}
		return true
	} else if a == nil && b == nil {
		return true
	}
	return false
}
