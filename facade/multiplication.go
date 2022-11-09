package facade

import (
	"constraint-system/core"
)

type multiplicationExpr struct {
	*baseExpr
}

func Multiplication(left, right []*core.Connector) Expr {
	expr := &multiplicationExpr{&baseExpr{
		left:     left,
		right:    right,
		variable: make(map[string]*core.Connector),
	}}
	expr.Bind()
	core.MultiplicationConstrain(expr.left, expr.right)
	return expr
}
