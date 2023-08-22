package sort_utils

import "math/rand"

func RandomListOfSize(size int) []int {
	out := make([]int, size)

	for i := 0; i < size; i++ {
		flip := rand.Intn(2)
		if flip == 0 {
			flip = -1
		}

		num := flip * rand.Intn(1000000)
		out[i] = num
	}

	return out
}
