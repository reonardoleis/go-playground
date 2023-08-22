package shell_sort

var ciuraSequence = []int{701, 301, 132, 57, 23, 10, 4, 1}

func ShellSort(numbers []int) ([]int, int) {
	iterations := 0
	for _, gap := range ciuraSequence {
		for i := gap; i < len(numbers); i++ {
			temp := numbers[i]

			j := i
			for ; (j >= gap) && (numbers[j-gap] > temp); j = j - gap {
				iterations++
				numbers[j] = numbers[j-gap]
			}

			numbers[j] = temp
		}
	}

	return numbers, iterations
}
