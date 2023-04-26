package classify

import (
	"constraint-system/facade"
	"math"
)

var (
	PythagoreanTheoremEdge1 = facade.Variable("edge1")
	PythagoreanTheoremEdge2 = facade.Variable("edge2")
	PythagoreanTheoremEdge3 = facade.Variable("edge3")
	PythagoreanTheorem      = facade.Addition(
		facade.Params(facade.Square(PythagoreanTheoremEdge1)),
		facade.Params(facade.Square(PythagoreanTheoremEdge2), facade.Square(PythagoreanTheoremEdge3)))
)

var (
	circleAreaExprRadius = facade.Variable("r")
	circleAreaExprArea   = facade.Variable("A")
	circleAreaExpr       = facade.Multiplication(
		facade.Params(circleAreaExprArea),
		facade.Params(facade.Constant(math.Pi), facade.Square(circleAreaExprRadius)))
)
