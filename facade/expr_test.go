package facade

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Complex_Calculate(t *testing.T) {
	Convey("乘法&&加法", t, func() {
		Convey("a = 5 * (b + 3)", func() {
			sumExpr := Addition(params(Variable("temp")), params(Variable("b"), Constant(3)))
			temp := sumExpr.GetVariable("temp")
			complexExpr := Multiplication(params(Variable("a")), params(Constant(5), temp))
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
	})
}
