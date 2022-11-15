package facade

import (
	"constraint-system/core"
)

func Square(origin *core.Connector) *core.Connector {
	newOne := Intermediate()
	core.SquareConstrain(newOne, origin)
	return newOne
}
