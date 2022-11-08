package core

func MultiplicationConstrain(left, right []*Connector) *constrain {
	return MakeTernaryConstraint(left, right, func(leftV, rightV []float64) float64 {
		var product float64 = 1
		for _, param := range rightV {
			product *= param
		}
		for _, param := range leftV {
			product /= param
		}
		return product
	}, func(leftV, rightV []float64) float64 {
		var product float64 = 1
		for _, param := range leftV {
			product *= param
		}
		for _, param := range rightV {
			product /= param
		}
		return product
	})
}
