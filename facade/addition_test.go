package facade

import (
	"constraint-system/core"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddition(t *testing.T) {
	Convey("加法", t, func() {
		Convey("a = b + 10", func() {
			expr := Addition(Params(Variable("a")), Params(Variable("b"), Constant(10)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			Convey("求a", func() {
				b.SetValue("user", 10)
				So(a.GetValue(), ShouldEqual, 20)
				So(b.GetValue(), ShouldEqual, 10)
			})
			Convey("求b", func() {
				a.SetValue("user", 20)
				So(a.GetValue(), ShouldEqual, 20)
				So(b.GetValue(), ShouldEqual, 10)
				a.ForgotValue("user")
				a.SetValue("user", 30)
				So(a.GetValue(), ShouldEqual, 30)
				So(b.GetValue(), ShouldEqual, 20)
			})

		})
		Convey("20 + a = b + 10", func() {
			expr := Addition(Params(Constant(20), Variable("a")), Params(Variable("b"), Constant(10)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			Convey("求a", func() {
				b.SetValue("user", 10)
				So(a.GetValue(), ShouldEqual, 0)
				So(b.GetValue(), ShouldEqual, 10)
			})
			Convey("求b", func() {
				a.SetValue("user", 20)
				So(a.GetValue(), ShouldEqual, 20)
				So(b.GetValue(), ShouldEqual, 30)
			})
		})
		Convey("计算：a = b + c + 10 + 5,条件不满足", func() {
			expr := Addition(Params(Variable("a")), Params(Variable("b"), Variable("c"), Constant(10), Constant(5)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			c := expr.GetVariable("c")
			b.SetValue("user", 5)
			defer func() {
				if err := recover(); err != nil {
					So(err, ShouldNotBeNil)
					So(err, ShouldEqual, core.NoValueErr)
				}
			}()
			a.GetValue()
			c.GetValue()
		})
	})
}
