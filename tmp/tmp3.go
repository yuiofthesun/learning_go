package main

import "fmt"

func main() {
	var testmap = map[string]int{
	"kat": 1,
	"mat": 23,
	"foo": 13,
	"bar": 34,
}
	i := len(testmap)
	_, ok := testmap["kat"]
	// := testmap["root"]

	for key, value := range testmap {
		fmt.Println( i, key, value, ok)
	}
}
