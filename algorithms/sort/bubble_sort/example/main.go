package main

import (
	"fmt"

	"github.com/reonardoleis/go-playground/algorithms/sort/bubble_sort"
)

func main() {
	sorted, iterations := bubble_sort.BubbleSort([]int{3, 2, 1, 0})
	fmt.Printf("%+v\n%d", sorted, iterations)
}
