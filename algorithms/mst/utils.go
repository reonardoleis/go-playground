package mst

import "math"

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getTableHeader(c int) string {
	out := ""

	i := int(math.Floor(float64(c) / float64(len(alphabet))))

	for ; i >= 0; i-- {
		curr := c - (len(alphabet) * i)

		pos := curr % len(alphabet)
		if curr >= len(alphabet) {
			sub := len(alphabet) * i
			pos = (c - sub) % int(math.Max(float64(i-1), 1))
		}

		out += string(alphabet[pos])
	}

	reversed := ""
	for i := len(out) - 1; i >= 0; i-- {
		reversed += string(out[i])
	}

	return reversed
}
