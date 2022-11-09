package facade

import (
	"constraint-system/core"
	"fmt"
	"github.com/google/uuid"
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
