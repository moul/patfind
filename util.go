package main

import (
	"sort"
)

func LeftSplitByLength(s string, length int) []string {
	var output []string
	var res string
	for i, r := range s {
		res = res + string(r)
		// fmt.Printf("i%d r %c\n", i, r)
		if i > 0 && (i+1)%length == 0 {
			output = append(output, res)
			// fmt.Printf("=>(%d) '%v'\n", i, res)
			res = ""
		}
	}
	return output
}

func RightSplitByLength(s string, length int) []string {
	return LeftSplitByLength(s[len(s)%length:], length)
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func StringUniqueSymbols(s string) []rune {
	output := []rune{}
	for _, c := range s {
		found := false
		for i := range output {
			if output[i] == c {
				found = true
				break
			}
		}
		if !found {
			output = append(output, c)
		}
	}
	sort.Sort(RuneSlice(output))
	return output
}
