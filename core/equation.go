package core

func EquationConstrain(left, right *Connector) *constrain {
	return MakeTernaryConstraint([]*Connector{left}, []*Connector{right}, func(leftV, rightV []float64) float64 {
		return right.GetValue()
	}, func(leftV, rightV []float64) float64 {
		return left.GetValue()
	})
}
