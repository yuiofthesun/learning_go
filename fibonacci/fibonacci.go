package main

import "fmt"

func fibonacci() func(int) int {
	start := 0
	start1 := 1
	return func(x int) int {
		var fib int =  start + start1
		start1 = start
		start = fib
		return fib
	}
}

func main() {
	f:= fibonacci()
	for i := 0; i<10; i++ {
		fmt.Println(f(i))
	}
}
