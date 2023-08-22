package bubble_sort

func BubbleSort(numbers []int) ([]int, int) {
	iterations := 0
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			iterations++
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	return numbers, iterations
}
