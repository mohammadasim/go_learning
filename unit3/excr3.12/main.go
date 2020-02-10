package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagramsCheck(a string, b string) int {
	// A function to check if the two imput strings are anagrams
	aString := strings.ToLower(a)
	bString := strings.ToLower(b)
	aSplitString := strings.Split(aString, "")
	bSplitString := strings.Split(bString, "")

	sort.Strings(aSplitString)
	sort.Strings(bSplitString)
	fmt.Println(aSplitString)
	fmt.Println(bSplitString)

	return strings.Compare(aString, bString)
}

func main() {
	anagramsCheck("adezxnobc", "cqrabegkipqde")
}
