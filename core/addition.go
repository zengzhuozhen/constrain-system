package core

func AdditionConstrain(left, right []*Connector) *constrain {
	return MakeTernaryConstraint(left, right, func(leftV, rightV []float64) float64 {
		var sum float64
		for _, param := range rightV {
			sum += param
		}
		for _, param := range leftV {
			sum -= param
		}
		return sum

	}, func(leftV, rightV []float64) float64 {
		var sum float64
		for _, param := range leftV {
			sum += param
		}
		for _, param := range rightV {
			sum -= param
		}
		return sum
	})
}
