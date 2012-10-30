package main

import "tour/pic"q
//import "fmt"

func Pic(dx, dy int) [][]uint8 {
	work := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		work[x] = make([]uint8, dy)
		for y,_ := range work[x] {
			work[x][y] = uint8(x^y)
		}
		
	}
	return work
}

func main() {
	pic.Show(Pic)
}
