package helpers

// Min finds the smallest of the values
func Min(n ...int) (min int) {
	min = int(^uint(0) >> 1)
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	return
}

// Max finds the largest of the values
func Max(n ...int) (max int) {
	max = -int(^uint(0)>>1) - 1
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return
}
