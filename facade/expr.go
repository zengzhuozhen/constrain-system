package facade

import (
	"constraint-system/core"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

var IntermediateConnectorName = "intermediate"

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

func Intermediate() *core.Connector {
	return core.MakeConnector(fmt.Sprintf("%s_%s", IntermediateConnectorName, uuid.New().String()))
}

type Expr interface {
	GetVariable(name string) *core.Connector
	GetIntermediate() *core.Connector
}

type baseExpr struct {
	left         []*core.Connector
	right        []*core.Connector
	variable     map[string]*core.Connector
	intermediate *core.Connector
}

func (a *baseExpr) GetVariable(name string) *core.Connector {
	return a.variable[name]
}

func (a *baseExpr) GetIntermediate() *core.Connector {
	return a.intermediate
}

func (a *baseExpr) Bind() {
	for _, i := range append(a.left, a.right...) {
		if i.Name != "" {
			a.variable[i.Name] = i
		}
		if strings.Contains(i.Name, IntermediateConnectorName) {
			a.intermediate = i
			a.intermediate.Name = ""
		}
	}
}
