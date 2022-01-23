package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	words := strings.Fields(s)
	
	wordMap := make(map[string]int)
	
	/* populate map */
	for i := range words { 
		word := words[i]
		elem, ok := wordMap[word]
		if ok { 
			/* word exists in the map */
			delete(wordMap,word)
			wordMap[word] = elem + 1
		} else { 
			/* word new to the map */
			wordMap[word] = 1
		}
	}
	
	return wordMap;
}

func main() {
	wc.Test(WordCount)
}
