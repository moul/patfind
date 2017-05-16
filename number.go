package main

import "fmt"

type Number string

func New(number string) Number {
	return Number(number)
}

func (n Number) String() string {
	return string(n)
}

func (n Number) UniqueSymbols() []rune {
	return StringUniqueSymbols(string(n))
}

func (n Number) RepetitivePatterns() map[string]int {
	fmt.Println("")
	fmt.Println(2, RightSplitByLength(string(n), 2))
	fmt.Println(3, RightSplitByLength(string(n), 3))
	fmt.Println(4, RightSplitByLength(string(n), 4))
	fmt.Println(5, RightSplitByLength(string(n), 5))
	return nil
}

func (n Number) Score() float64 {
	// FIXME: check for famous patterns, 4242, 666
	// FIXME: check for suites
	return 42
}
