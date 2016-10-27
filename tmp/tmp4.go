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
	for  i:=0; i < 11; i++ {
		ReturnMap[splitWords[i]]= wordcount
		fmt.Println(ReturnMap)
//		for j := i; j<NumberofWords; j++{
//			firstTerm := splitWords[i]
//			CompareTerm := splitWords[j]
//			if firstTerm == CompareTerm {
//				wordcount++
//			}
//		}
	}
	fmt.Println(ReturnMap)
}


