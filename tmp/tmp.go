package main

import "fmt"

func main() {

	ReturnSlice := make([][]uint8,dy)
		for i:=0; i<dy; dy++ {
			InnerLen := dy +1
			ReturnSlice[dy] =  make([]unit8, InnerLen)
			for j :=0; j< InnerLen; j++ {
				ReturnSlice[dy][j] = dy + j
			}
		}
	fmt.Println(sl1)


}
