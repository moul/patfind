package main

import (
	"fmt"
	"strings"
)

type Handler func(string) (float64, string)

var EnabledHandlers = []Handler{
	UniqueSymbolsScore,
	RepetitivePatternsScore,
	FamousPatternsScore,
	// FIXME: check for suites
	// FIXME: longest ending similar numbers (/000$/)
	// FIXME: finger-typing-distance
	// FIXME: longest repetitive symbol
}

var FamousPatterns = []string{
	"42",
	"1337",
	"8173",
}

func FamousPatternsScore(n string) (score float64, comment string) {
	matchedPatterns := []string{}
	for _, pattern := range FamousPatterns {
		if strings.Contains(n, pattern) {
			matchedPatterns = append(matchedPatterns, pattern)
			score += 0.25
		}
		if strings.HasSuffix(n, pattern) {
			score += 0.5
		}
	}

	if score > 1 {
		score = 1
	}
	if score > 0 {
		comment = fmt.Sprintf("famous patterns: %s", strings.Join(matchedPatterns, ","))
	}
	return
}

func UniqueSymbolsScore(n string) (score float64, comment string) {
	uniq := StringUniqueSymbols(n)
	if len(uniq) > len(n)/2 {
		return
	}
	comment = fmt.Sprintf("%d unique symbols", len(uniq))
	uniqueness := float64(len(n)) / (float64(len(uniq)) / (float64(len(n)))) / 100 // average percentage of presence per symbol
	// FIXME: use an exponential approach over the uniqueness
	score = uniqueness
	return
}

func RepetitivePatternsScore(n string) (score float64, comment string) {
	var (
		bestMatchedLen int
		// FIXME: count the number of working "i"
	)

	for i := len(n) / 2; i >= 2; i-- {
		groups := RightSplitByLength(string(n), i)
		uniq := map[string]int{}
		for _, group := range groups {
			if _, found := uniq[group]; !found {
				uniq[group] = 0
			}
			uniq[group]++
		}

		matchedLen := 0
		sumMatches := 0
		for idx, val := range uniq {
			if val == 1 {
				delete(uniq, idx)
			} else {
				matchedLen += i * val
				sumMatches += val
			}
		}
		if matchedLen > bestMatchedLen {
			bestMatchedLen = matchedLen
			comment = fmt.Sprintf("%d repetitions of %d symbols", sumMatches, i)
		}
	}

	score = float64(bestMatchedLen) / float64(len(n))
	return
}
