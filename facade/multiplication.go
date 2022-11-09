package facade

import (
	"constraint-system/core"
	"strings"
)

type multiplicationExpr struct {
	left         []*core.Connector
	right        []*core.Connector
	variable     map[string]*core.Connector
	intermediate *core.Connector
}

func (a *multiplicationExpr) GetVariable(name string) *core.Connector {
	return a.variable[name]
}

func (a *multiplicationExpr) GetIntermediate() *core.Connector {
	return a.intermediate
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
		if strings.Contains(i.Name, IntermediateConnectorName) {
			expr.intermediate = i
			expr.intermediate.Name = ""
		}
	}
	core.MultiplicationConstrain(expr.left, expr.right)
	return expr
}
