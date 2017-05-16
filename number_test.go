package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumber_Score(t *testing.T) {
	Convey("Number.Score", t, func() {
		So(New("0123456789").Score(), ShouldEqual, 42)
		So(New("0958274251").Score(), ShouldEqual, 42)
		So(New("0666666666").Score(), ShouldEqual, 42)
		So(New("0655554444").Score(), ShouldEqual, 42)
		So(New("0651510000").Score(), ShouldEqual, 42)
		So(New("0652255225").Score(), ShouldEqual, 42)
	})
}
