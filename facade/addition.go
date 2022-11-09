package facade

import (
	"constraint-system/core"
)

type additionExpr struct {
	*baseExpr
}

func Addition(left, right []*core.Connector) Expr {
	expr := &additionExpr{&baseExpr{
		left:     left,
		right:    right,
		variable: make(map[string]*core.Connector),
	}}
	expr.Bind()
	core.AdditionConstrain(expr.left, expr.right)
	return expr
}
