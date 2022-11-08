package facade

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEquation(t *testing.T) {
	Convey("等值", t, func() {
		Convey("a = b", func() {
			expr := Equation(Variable("a"), Variable("b"))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			Convey("求a", func() {
				b.SetValue("user", 10)
				So(a.GetValue(), ShouldEqual, 10)
				So(b.GetValue(), ShouldEqual, 10)
			})
			Convey("求b", func() {
				a.SetValue("user", 10)
				So(a.GetValue(), ShouldEqual, 10)
				So(b.GetValue(), ShouldEqual, 10)
			})
		})
	})
}
