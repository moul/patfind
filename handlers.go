package main

import (
	"fmt"
)

type Handler func(string) (float64, string)

var EnabledHandlers = []Handler{
	UniqueSymbols,
	// RepetitivePatterns
	// FIXME: check for famous patterns, 4242, 666
	// FIXME: check for suites
}

func UniqueSymbols(n string) (score float64, comment string) {
	uniq := StringUniqueSymbols(n)
	comment = fmt.Sprintf("%d unique symbols", len(uniq))
	uniqueness := float64(len(n)) / (float64(len(uniq)) / (float64(len(n)))) / 100 // average percentage of presence per symbol
	// FIXME: use an exponential approach over the uniqueness
	score = uniqueness
	return
}

func RepetitivePatterns(n string) (score float64, comment string) {
	fmt.Println("")
	fmt.Println(2, RightSplitByLength(string(n), 2))
	fmt.Println(3, RightSplitByLength(string(n), 3))
	fmt.Println(4, RightSplitByLength(string(n), 4))
	fmt.Println(5, RightSplitByLength(string(n), 5))
	return
}
