package main

import "testing"

func grid(rows ...string) [][]byte {
	g := make([][]byte, len(rows))
	for i, r := range rows {
		g[i] = []byte(r)
	}
	return g
}

func TestNumIslands(t *testing.T) {
	cases := []struct {
		g    [][]byte
		want int
	}{
		{grid("11000", "11000", "00100", "00011"), 3},
		{grid("111", "010", "111"), 1},
		{grid("101", "010", "101"), 5}, // diagonals don't connect
		{grid("000"), 0},
		{grid(), 0},
		{grid("1"), 1},
	}
	for _, c := range cases {
		if got := NumIslands(c.g); got != c.want {
			t.Errorf("NumIslands = %d, want %d", got, c.want)
		}
	}
}
