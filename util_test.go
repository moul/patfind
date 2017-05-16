package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLeftSplitByLength(t *testing.T) {
	Convey("LeftSplitByLength", t, func() {
		So(LeftSplitByLength("0123456789", 2), ShouldResemble, []string{"01", "23", "45", "67", "89"})
		So(LeftSplitByLength("0123456789", 3), ShouldResemble, []string{"012", "345", "678"})
		So(LeftSplitByLength("0123456789", 4), ShouldResemble, []string{"0123", "4567"})
		So(LeftSplitByLength("0123456789", 5), ShouldResemble, []string{"01234", "56789"})
	})
}

func TestRightSplitByLength(t *testing.T) {
	Convey("RightSplitByLength", t, func() {
		So(RightSplitByLength("0123456789", 2), ShouldResemble, []string{"01", "23", "45", "67", "89"})
		So(RightSplitByLength("0123456789", 3), ShouldResemble, []string{"123", "456", "789"})
		So(RightSplitByLength("0123456789", 4), ShouldResemble, []string{"2345", "6789"})
		So(RightSplitByLength("0123456789", 5), ShouldResemble, []string{"01234", "56789"})
	})
}

func TestStringUniqueSymbols(t *testing.T) {
	Convey("StringUniqueSymbols", t, func() {
		So(StringUniqueSymbols("0777777777"), ShouldResemble, []rune{'0', '7'})
		So(StringUniqueSymbols("0123456789"), ShouldResemble, []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})
		So(StringUniqueSymbols("8651302479"), ShouldResemble, []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})
	})
}
