package main

import (
	"fmt"
	"sort"
	"unicode"
)

func arePermutations(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	sorted1 := sortString(s1)
	sorted2 := sortString(s2)
	return sorted1 == sorted2
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func main() {
	fmt.Println(arePermutations("abc", "cbd"))
	unicode.IsLower(rune('a'))
}
