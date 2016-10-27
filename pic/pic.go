package main

import "golang.org/x/tour/pic"


func makePic(formula func (x,y int) uint8) func (dx,dy int)[][]uint8 {
	return func (dx,dy int)[][]uint8{
		returnSlice := make([][]uint8,dy)
		for y := range returnSlice {
			returnSlice[y] = make([]uint8, dx)
			for x:= range returnSlice[y] {
				returnSlice[y][x] = formula(x, y)
			}
		}
		return returnSlice
	}
}

func main() {

	formula1 := func (x, y int) uint8{
		return uint8((x+y)/2)
	 }
	formula2 := func (x, y int) uint8{
		return uint8((x*y))
	 }
	pic.Show(makePic(formula1))
	pic.Show(makePic(formula2))
}

