package facade

import "constraint-system/core"

func Constant(v float64) *core.Connector {
	connector := core.MakeConnector("")
	connector.SetValue("", v)
	return connector
}

func Variable(name string) *core.Connector {
	connector := core.MakeConnector(name)
	return connector
}

func params(c ...*core.Connector) (cc []*core.Connector) {
	return append(cc, c...)
}

type (
	LeftExpr  func() Expr
	RightExpr func() Expr
)

type Expr interface {
	GetVariable(name string) *core.Connector
}

type additionExpr struct {
	left     []*core.Connector
	right    []*core.Connector
	variable map[string]*core.Connector
}

func (a *additionExpr) GetVariable(name string) *core.Connector {
	return a.variable[name]
}

func Addition(left, right []*core.Connector) Expr {
	expr := &additionExpr{
		left:     left,
		right:    right,
		variable: make(map[string]*core.Connector),
	}
	for _, i := range append(expr.left, expr.right...) {
		if i.Name != "" {
			expr.variable[i.Name] = i
		}
	}
	core.AdditionConstrain(expr.left, expr.right)
	return expr
}
