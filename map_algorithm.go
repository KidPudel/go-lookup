package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for _, word := range words {
		_, alreadyExists := wordCount[word]
		if alreadyExists {
			wordCount[word]++
		} else {
			wordCount[word] = 1
			
			a := 3
			
		
		}
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
