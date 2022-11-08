package facade

import (
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
			sumExpr := Addition(Params(Variable("tempSum")), Params(Variable("b"), Constant(4)))
			b := sumExpr.GetVariable("b")
			multiExpr := Multiplication(Params(Variable("tempMulti")), Params(Constant(2), sumExpr.GetVariable("tempSum")))
			leftMultiExpr := Multiplication(Params(Constant(2), Variable("a")), Params(Variable("tempMultiLeft")))
			a := leftMultiExpr.GetVariable("a")
			Addition(Params(leftMultiExpr.GetVariable("tempMultiLeft"), Constant(4)),
				Params(multiExpr.GetVariable("tempMulti")))
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
	})
}
