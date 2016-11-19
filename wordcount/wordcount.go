package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount (s string) map[string]int {
	ReturnMap := make(map[string]int)
	splitWords := strings.Fields(s)
	NumberofWords:= len(splitWords)
	var wordcount int
	for , splitWords[i] := range splitWords {
		ReturnMap[splitWords[i]]= wordcount

	//j := i
//	for j, splitWords[j] := range ReturnMap {
//		firstTerm := splitWords[i]
//		CompareTerm := splitWords[j]
//			if firstTerm == CompareTerm {
//				wordcount++
			}
}
	return ReturnMap
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
