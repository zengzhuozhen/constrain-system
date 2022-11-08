package facade

import (
	"constraint-system/core"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddition(t *testing.T) {
	Convey("加法", t, func() {
		Convey("计算：a = b + 10,求a", func() {
			expr := Addition(params(Variable("a")), params(Variable("b"), Constant(10)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			b.SetValue("user", 10)
			So(a.GetValue(), ShouldEqual, 20)
			So(b.GetValue(), ShouldEqual, 10)
		})
		Convey("计算：20 + a = b + 10,求a", func() {
			expr := Addition(params(Constant(20), Variable("a")), params(Variable("b"), Constant(10)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			b.SetValue("user", 10)
			So(a.GetValue(), ShouldEqual, 0)
			So(b.GetValue(), ShouldEqual, 10)
		})
		Convey("计算：20 + a = b + 10,求b", func() {
			expr := Addition(params(Constant(20), Variable("a")), params(Variable("b"), Constant(10)))
			a := expr.GetVariable("a")
			b := expr.GetVariable("b")
			a.SetValue("user", 10)
			So(a.GetValue(), ShouldEqual, 10)
			So(b.GetValue(), ShouldEqual, 20)
		})
		Convey("计算：a = b + c + 10 + 5,求a,条件不满足", func() {
			expr := Addition(params(Variable("a")), params(Variable("b"), Variable("c"), Constant(10), Constant(5)))
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
