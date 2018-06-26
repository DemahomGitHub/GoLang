package main

import (
	"fmt"
)

var terms = [20]string{
	"new", "home", "sales", "top", "forecasts",
	"home", "sales", "rise", "in", "july",
	"increase", "in", "home", "sales", "july",
	"july", "new", "home", "sales", "rise",
}

func main() {
	termsMap := make(map[string]int)

	for i, key := range terms[1:] {
		j := i - 1
		for j >= 0 && key < terms[j] {
			terms[j+1] = terms[j]
			j--
		}
		terms[j+1] = key
		termsMap[key]++
	}
	fmt.Println(terms)
	printTerms(termsMap)
}

func printTerms(terms map[string]int) {
	for id, val := range terms {
		fmt.Println(id+" : ", val)
	}
}
