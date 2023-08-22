package main

import (
	"fmt"
	"time"

	"github.com/reonardoleis/go-playground/algorithms/sort/bubble_sort"
	"github.com/reonardoleis/go-playground/algorithms/sort/shell_sort"
	sort_utils "github.com/reonardoleis/go-playground/algorithms/sort/utils"
)

type sortingAlgorithm struct {
	name      string
	algorithm func(numbers []int) (sorted []int, iterations int)
}

var algorithms []sortingAlgorithm = []sortingAlgorithm{
	{"bubble sort", bubble_sort.BubbleSort},
	{"shell sort", shell_sort.ShellSort},
}

func main() {
	list := sort_utils.RandomListOfSize(10000)

	for _, algorithm := range algorithms {
		listCopy := make([]int, len(list))
		copy(listCopy, list)

		start := time.Now()

		_, iterations := algorithm.algorithm(listCopy)

		end := time.Now()

		fmt.Printf("Algorithm: %s\nIterations: %d\nTime (in ms): %d\nTime (in μs): %d\n\n--\n\n",
			algorithm.name,
			iterations,
			end.Sub(start).Milliseconds(),
			end.Sub(start).Microseconds())
	}
}
