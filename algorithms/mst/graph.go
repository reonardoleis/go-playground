package mst

import "fmt"

type Graph struct {
	table [][]int
}

func NewGraph(table [][]int) Graph {
	return Graph{table}
}

func (g Graph) String() string {
	out := ""
	for j := 0; j < len(g.table[0]); j++ {
		out += ("\t" + getTableHeader(j))
	}

	for i := 0; i < len(g.table); i++ {
		out += ("\n" + getTableHeader(i) + "\t")
		for j := 0; j < len(g.table[i]); j++ {
			out += fmt.Sprintf("%d\t", g.table[i][j])
		}

	}

	return out
}

func (g Graph) PrimMST() Graph {
	// TODO
	return g
}
