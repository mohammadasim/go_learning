package main

import (
	"sort"
	"strings"
)

func anagramsCheck(a string, b string) bool {
	// A function to check if the two imput strings are anagrams
	// change both strings to lower case
	aString := strings.ToLower(a)
	bString := strings.ToLower(b)

	// Change both strings to string slices
	aSplitString := strings.Split(aString, "")
	bSplitString := strings.Split(bString, "")

	//Sort both string slices
	sort.Strings(aSplitString)
	sort.Strings(bSplitString)

	//change back the sorted string slices to strings
	aSortedString := strings.Join(aSplitString, ",")
	bSortedString := strings.Join(bSplitString, ",")

	return aSortedString == bSortedString
}

func main() {
	println(anagramsCheck("abcd", "bca"))
}
