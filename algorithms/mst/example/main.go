package main

import "github.com/reonardoleis/go-playground/algorithms/mst"

func main() {
	g := mst.NewGraph(
		[][]int{
			{1, 2, 3},
			{0, 2, 3},
		},
	)

	println(g.String())

	mst := g.PrimMST()

	println(mst.String())
}
