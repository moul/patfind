package main

import (
	"fmt"
)

type Handler func(string) (float64, string)

var EnabledHandlers = []Handler{
	UniqueSymbols,
	RepetitivePatterns,
	// FIXME: check for famous patterns, 4242, 666
	// FIXME: check for suites
	// FIXME: longest ending similar numbers (/000$/)
	// FIXME: finger-typing-distance
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
	var (
		bestMatchedLen int
	)

	// FIXME: begin at i:=1 ?
	for i := 2; i < len(n)/2; i++ {
		groups := RightSplitByLength(string(n), i)
		uniq := map[string]int{}
		for _, group := range groups {
			if _, found := uniq[group]; !found {
				uniq[group] = 0
			}
			uniq[group]++
		}

		matchedLen := 0
		for idx, val := range uniq {
			if val == 1 {
				delete(uniq, idx)
			} else {
				matchedLen += i * val
			}
		}
		if matchedLen > bestMatchedLen {
			bestMatchedLen = matchedLen
		}
	}

	score = float64(bestMatchedLen) / float64(len(n))
	return
}
