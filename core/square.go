package core

import "math"

func SquareConstrain(left, right *Connector) *constrain {
	return MakeTernaryConstraint([]*Connector{left}, []*Connector{right}, func(leftV, rightV []float64) float64 {
		return right.GetValue() * right.GetValue()
	}, func(leftV, rightV []float64) float64 {
		return math.Sqrt(left.GetValue())
	})
}
