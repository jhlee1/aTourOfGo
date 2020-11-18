package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)

	for x := 0; x < dy; x++ {
		b := make([]uint8, dx)

		for y := range b {
			b[y] = uint8((x + y) / 2)
		}

		picture[x] = b
	}

	return picture

}

func main() {
		pic.Show(Pic)
}
