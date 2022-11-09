package facade

import (
	"constraint-system/core"
	"strings"
)

type equationExpr struct {
	left         *core.Connector
	right        *core.Connector
	variable     map[string]*core.Connector
	intermediate *core.Connector
}

func (a *equationExpr) GetVariable(name string) *core.Connector {
	return a.variable[name]
}

func (a *equationExpr) GetIntermediate() *core.Connector {
	return a.intermediate
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
		if strings.Contains(i.Name, IntermediateConnectorName) {
			expr.intermediate = i
			expr.intermediate.Name = ""
		}
	}
	core.EquationConstrain(left, right)
	return expr
}
