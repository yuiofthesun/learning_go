package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world here is kat Because a world kat is here"
	splitWords := strings.Fields(s)
	ReturnMap := make(map[string]int)
	NumberofWords :=len(splitWords)
	fmt.Println(NumberofWords)
	fmt.Println(splitWords)
	var wordcount int = 1
//for key, value := range m {
//    fmt.Println("Key:", key, "Value:", value)
//}  according to blog.golang.org/go-maps-in-action
//key hier ist jeweilige string, value meine Integer, richtig?
	for  splitWords[i], i:=0 := range ReturnMap {
		fmt.Println(splitWords[i])
		ReturnMap[splitWords[i]]= wordcount
		fmt.Println(ReturnMap)
	}
}
