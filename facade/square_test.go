package facade

import (
	"constraint-system/core"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSquare(t *testing.T) {
	Convey("平方", t, func() {
		Convey("a = (2b)^2", func() {
			a := Variable("a")
			b := Variable("b")
			Equation(a, Square(
				Multiplication(Params(Intermediate()), Params(Constant(2), b)).GetIntermediate(),
			))
			Convey("求a", func() {
				b.SetValue(core.PerDefineUserSource, 10)
				So(a.GetValue(), ShouldEqual, 400)
				So(b.GetValue(), ShouldEqual, 10)
			})
			Convey("求b", func() {
				a.SetValue(core.PerDefineUserSource, 100)
				So(a.GetValue(), ShouldEqual, 100)
				So(b.GetValue(), ShouldEqual, 5)
			})
		})
		Convey("a^2 = 2b", func() {
			a, b := Variable("a"), Variable("b")
			Equation(Square(a),
				Multiplication(Params(Intermediate()), Params(Constant(2), b)).GetIntermediate())
			Convey("求a", func() {
				b.SetValue(core.PerDefineUserSource, 18)
				So(a.GetValue(), ShouldEqual, 6)
				So(b.GetValue(), ShouldEqual, 18)
			})
			Convey("求b", func() {
				a.SetValue(core.PerDefineUserSource, 10)
				So(a.GetValue(), ShouldEqual, 10)
				So(b.GetValue(), ShouldEqual, 50)
			})
		})
	})
}
