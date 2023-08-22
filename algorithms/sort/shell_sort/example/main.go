package main

import (
	"fmt"

	"github.com/reonardoleis/go-playground/algorithms/sort/shell_sort"
)

func main() {
	sorted, iterations := shell_sort.ShellSort([]int{3, 5, 2, 10, -1, 23, 23, 15})
	fmt.Printf("%+v\n%d", sorted, iterations)
}
