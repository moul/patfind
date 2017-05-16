package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueSymbols(t *testing.T) {
	Convey("UniqueSymbols", t, FailureContinues, func() {
		var (
			score   float64
			comment string
		)

		score, comment = UniqueSymbols("0000000000")
		So(score, ShouldEqual, float64(1))
		So(comment, ShouldEqual, "1 unique symbols")

		score, comment = UniqueSymbols("0707070707")
		So(score, ShouldEqual, float64(0.5))
		So(comment, ShouldEqual, "2 unique symbols")

		score, comment = UniqueSymbols("0123456789")
		So(score, ShouldEqual, float64(0.1))
		So(comment, ShouldEqual, "10 unique symbols")
	})
}

func TestRepetitivePatterns(t *testing.T) {
	Convey("RepetitivePatterns", t, FailureContinues, func() {
		var (
			score   float64
			comment string
		)

		score, comment = RepetitivePatterns("0000000000")
		So(score, ShouldEqual, float64(1))
		So(comment, ShouldEqual, "")

		score, comment = RepetitivePatterns("0707070707")
		So(score, ShouldEqual, float64(1))
		So(comment, ShouldEqual, "")

		score, comment = RepetitivePatterns("0682824242")
		So(score, ShouldEqual, 0.8)
		So(comment, ShouldEqual, "")

		score, comment = RepetitivePatterns("0612341234")
		So(score, ShouldEqual, 0.8)
		So(comment, ShouldEqual, "")

		score, comment = RepetitivePatterns("0612344242")
		So(score, ShouldEqual, 0.4)
		So(comment, ShouldEqual, "")

		score, comment = RepetitivePatterns("0612345666")
		So(score, ShouldEqual, 0)
		So(comment, ShouldEqual, "")

		score, comment = RepetitivePatterns("0123456789")
		So(score, ShouldEqual, float64(0))
		So(comment, ShouldEqual, "")
	})
}
