package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueSymbolsScore(t *testing.T) {
	Convey("UniqueSymbolsScore", t, FailureContinues, func() {
		var (
			score   float64
			comment string
		)

		score, comment = UniqueSymbolsScore("0000000000")
		So(score, ShouldEqual, float64(1))
		So(comment, ShouldEqual, "1 unique symbols")

		score, comment = UniqueSymbolsScore("0707070707")
		So(score, ShouldEqual, float64(0.5))
		So(comment, ShouldEqual, "2 unique symbols")

		score, comment = UniqueSymbolsScore("0123456789")
		So(score, ShouldEqual, float64(0.1))
		So(comment, ShouldEqual, "10 unique symbols")
	})
}

func TestFamousPatternsScore(t *testing.T) {
	Convey("FamousPatternsScore", t, FailureContinues, func() {
		var (
			score   float64
			comment string
		)

		score, comment = FamousPatternsScore("0606060606")
		So(score, ShouldEqual, float64(0))
		So(comment, ShouldEqual, "")

		score, comment = FamousPatternsScore("0619181742")
		So(score, ShouldEqual, float64(0.75))
		So(comment, ShouldEqual, "famous patterns: 42")

		score, comment = FamousPatternsScore("0619181337")
		So(score, ShouldEqual, float64(0.75))
		So(comment, ShouldEqual, "famous patterns: 1337")

		score, comment = FamousPatternsScore("0619184200")
		So(score, ShouldEqual, float64(0.25))
		So(comment, ShouldEqual, "famous patterns: 42")

		score, comment = FamousPatternsScore("0619133700")
		So(score, ShouldEqual, float64(0.25))
		So(comment, ShouldEqual, "famous patterns: 1337")

		score, comment = FamousPatternsScore("0642421234")
		So(score, ShouldEqual, float64(0.25))
		So(comment, ShouldEqual, "famous patterns: 42")

		score, comment = FamousPatternsScore("0642133712")
		So(score, ShouldEqual, float64(0.5))
		So(comment, ShouldEqual, "famous patterns: 42,1337")

		score, comment = FamousPatternsScore("0642421337")
		So(score, ShouldEqual, float64(1))
		So(comment, ShouldEqual, "famous patterns: 42,1337")
	})
}

func TestRepetitivePatternsScore(t *testing.T) {
	Convey("RepetitivePatternsScore", t, FailureContinues, func() {
		var (
			score   float64
			comment string
		)

		score, comment = RepetitivePatternsScore("0000000000")
		So(score, ShouldEqual, float64(1))
		So(comment, ShouldEqual, "2 repetitions of 5 symbols")

		score, comment = RepetitivePatternsScore("0707070707")
		So(score, ShouldEqual, float64(1))
		So(comment, ShouldEqual, "5 repetitions of 2 symbols")

		score, comment = RepetitivePatternsScore("0682824242")
		So(score, ShouldEqual, 0.8)
		So(comment, ShouldEqual, "4 repetitions of 2 symbols")

		score, comment = RepetitivePatternsScore("0612341234")
		So(score, ShouldEqual, 0.8)
		So(comment, ShouldEqual, "2 repetitions of 4 symbols")

		score, comment = RepetitivePatternsScore("0612344242")
		So(score, ShouldEqual, 0.4)
		So(comment, ShouldEqual, "2 repetitions of 2 symbols")

		score, comment = RepetitivePatternsScore("0612345666")
		So(score, ShouldEqual, 0)
		So(comment, ShouldEqual, "")

		score, comment = RepetitivePatternsScore("0123456789")
		So(score, ShouldEqual, float64(0))
		So(comment, ShouldEqual, "")
	})
}
