package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		a := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			switch x % 3 {
			case 1:
				a[x] = uint8((x + y) / 2)
			case 2:
				a[x] = uint8(x * y)
			case 3:
				a[x] = uint8(x ^ y)
			}
		}
		picture[y] = a
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
