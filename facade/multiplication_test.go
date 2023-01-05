package facade

import (
	"constraint-system/core"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMultiplication(t *testing.T) {
	Convey("乘法", t, func() {
		Convey("a = 9 * b", func() {
			expr := Multiplication(Params(Variable("a")), Params(Variable("b"), Constant(9)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			Convey("求a", func() {
				b.SetValue(core.PredefineUserSource, 10)
				So(a.GetValue(), ShouldEqual, 90)
				So(b.GetValue(), ShouldEqual, 10)
			})
			Convey("求b", func() {
				a.SetValue(core.PredefineUserSource, 45)
				So(a.GetValue(), ShouldEqual, 45)
				So(b.GetValue(), ShouldEqual, 5)
			})
		})
		Convey("5 * a = 20 * b", func() {
			expr := Multiplication(Params(Constant(5), Variable("a")), Params(Variable("b"), Constant(20)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			Convey("求a", func() {
				b.SetValue(core.PredefineUserSource, 1)
				So(a.GetValue(), ShouldEqual, 4)
				So(b.GetValue(), ShouldEqual, 1)
			})
			Convey("求b", func() {
				a.SetValue(core.PredefineUserSource, 10)
				So(a.GetValue(), ShouldEqual, 10)
				So(b.GetValue(), ShouldEqual, 2.5)
			})
		})
	})
}
