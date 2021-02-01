package main

import (
	"fmt"
	"sort"
)

type byLength []string

func main() {
	sortingSlice := []string{"this", "is", "a", "sorting", "test"}
	sort.Sort(byLength(sortingSlice))
	fmt.Println(sortingSlice)
}

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	fmt.Println(s)
	return len(s[i]) < len(s[j])
}
