package facade

import (
	"constraint-system/core"
)

func Constant(v float64) *core.Connector {
	connector := core.MakeConnector("")
	connector.SetValue("", v)
	return connector
}

func Variable(name string) *core.Connector {
	connector := core.MakeConnector(name)
	return connector
}

func Params(c ...*core.Connector) (cc []*core.Connector) {
	return append(cc, c...)
}

type (
	LeftExpr  func() Expr
	RightExpr func() Expr
)

type Expr interface {
	GetVariable(name string) *core.Connector
}
