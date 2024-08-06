package core

import "math"

func SquareConstrain(left, right *Connector) *constrain {
	return MakeTernaryConstraint([]*Connector{left}, []*Connector{right}, func(leftV, rightV []float64) float64 {
		return right.innerGet() * right.innerGet()
	}, func(leftV, rightV []float64) float64 {
		return math.Sqrt(left.innerGet())
	})
}
