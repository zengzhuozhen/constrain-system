package facade

import "constraint-system/core"

type equationExpr struct {
	left     *core.Connector
	right    *core.Connector
	variable map[string]*core.Connector
}

func (a *equationExpr) GetVariable(name string) *core.Connector {
	return a.variable[name]
}

func Equation(left, right *core.Connector) *equationExpr {
	expr := &equationExpr{
		left:     left,
		right:    right,
		variable: make(map[string]*core.Connector),
	}
	for _, i := range append([]*core.Connector{left}, right) {
		if i.Name != "" {
			expr.variable[i.Name] = i
		}
	}
	core.EquationConstrain(left, right)
	return expr
}
