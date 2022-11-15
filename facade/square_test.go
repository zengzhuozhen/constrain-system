package facade

import (
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
				b.SetValue("user", 10)
				So(a.GetValue(), ShouldEqual, 400)
				So(b.GetValue(), ShouldEqual, 10)
			})
			Convey("求b", func() {
				a.SetValue("user", 100)
				So(a.GetValue(), ShouldEqual, 100)
				So(b.GetValue(), ShouldEqual, 5)
			})
		})
	})
}
