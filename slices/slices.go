package slices

import "golang.org/x/tour/pic"

type SliceDemo struct{}

func (s SliceDemo) Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			result[i][j] = uint8((i + j) / 2)
		}
	}
	return result
}

func (s SliceDemo) Execute() {
	pic.Show(s.Pic)
}
