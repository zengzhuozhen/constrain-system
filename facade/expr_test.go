package facade

import (
	"constraint-system/core"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Complex_Calculate(t *testing.T) {
	Convey("乘法&&加法", t, func() {
		Convey("a = 5 * (b + 3)", func() {
			sumExpr := Addition(Params(Variable("temp")), Params(Variable("b"), Constant(3)))
			complexExpr := Multiplication(Params(Variable("a")), Params(Constant(5), sumExpr.GetVariable("temp")))
			a := complexExpr.GetVariable("a")
			b := sumExpr.GetVariable("b")
			Convey("求a", func() {
				b.SetValue("user", 7)
				So(a.GetValue(), ShouldEqual, 50)
				So(b.GetValue(), ShouldEqual, 7)
			})
			Convey("求b", func() {
				a.SetValue("user", 100)
				So(a.GetValue(), ShouldEqual, 100)
				So(b.GetValue(), ShouldEqual, 17)
			})
		})

		Convey("(2 * a) + 4 = 2 * (b + 4)", func() {
			b := Variable("b")
			a := Variable("a")
			Equation(
				// left
				Addition(Params(Intermediate()), Params(
					Multiplication(Params(Intermediate()), Params(Constant(2), a)).GetIntermediate(), Constant(4)),
				).GetIntermediate(),
				// right
				Multiplication(Params(Intermediate()), Params(Constant(2),
					Addition(Params(Intermediate()), Params(b, Constant(4))).GetIntermediate()),
				).GetIntermediate())
			Convey("求a", func() {
				b.SetValue("user", 4)
				So(a.GetValue(), ShouldEqual, 6)
				So(b.GetValue(), ShouldEqual, 4)
			})
			Convey("求b", func() {
				a.SetValue("user", 4)
				So(a.GetValue(), ShouldEqual, 4)
				So(b.GetValue(), ShouldEqual, 2)
			})
		})

		Convey("9C = 5(F-32)", func() {
			C, F := Variable("c"), Variable("f")
			Multiplication(Params(Constant(9), C), Params(Constant(5),
				Addition(Params(Intermediate()), Params(F, Constant(-32))).GetIntermediate()))
			Convey("求c", func() {
				F.SetValue("user", 50)
				So(C.GetValue(), ShouldEqual, 10)
				So(F.GetValue(), ShouldEqual, 50)
			})
			Convey("求f", func() {
				C.SetValue("user", 10)
				So(C.GetValue(), ShouldEqual, 10)
				So(F.GetValue(), ShouldEqual, 50)
			})
		})

		Convey("a = 2b + b^2", func() {
			a, b := Variable("a"), Variable("b")
			Addition(Params(a), Params(
				Multiplication(Params(Intermediate()), Params(Constant(2), b)).GetIntermediate(), Square(b)))
			Convey("求a", func() {
				b.SetValue("user", 4)
				So(a.GetValue(), ShouldEqual, 24)
				So(b.GetValue(), ShouldEqual, 4)
			})
			Convey("求b,无法求解,有不同解", func() {
				defer func() {
					if err := recover(); err != nil {
						So(err, ShouldEqual, core.NoValueErr)
					}
				}()
				a.SetValue("user", 24)
				So(a.GetValue(), ShouldEqual, 24)
				So(b.GetValue(), ShouldEqual, 4)
			})
		})
	})
}
