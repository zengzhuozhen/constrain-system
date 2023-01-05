package facade

import (
	"constraint-system/core"
)

type equationExpr struct {
	*baseExpr
}

func (a *equationExpr) GetVariable(name string) *core.Connector {
	return a.variable[name]
}

func (a *equationExpr) GetIntermediate() *core.Connector {
	return a.intermediate
}

func Equation(left, right *core.Connector) Expr {
	expr := &equationExpr{&baseExpr{
		left:     []*core.Connector{left},
		right:    []*core.Connector{right},
		variable: make(map[string]*core.Connector),
	}}
	expr.Bind()
	core.EquationConstrain(left, right)
	return expr
}
