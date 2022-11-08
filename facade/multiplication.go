package facade

import "constraint-system/core"

type multiplicationExpr struct {
	left     []*core.Connector
	right    []*core.Connector
	variable map[string]*core.Connector
}

func (a *multiplicationExpr) GetVariable(name string) *core.Connector {
	return a.variable[name]
}

func Multiplication(left, right []*core.Connector) Expr {
	expr := &multiplicationExpr{
		left:     left,
		right:    right,
		variable: make(map[string]*core.Connector),
	}
	for _, i := range append(expr.left, expr.right...) {
		if i.Name != "" {
			expr.variable[i.Name] = i
		}
	}
	core.MultiplicationConstrain(expr.left, expr.right)
	return expr
}
