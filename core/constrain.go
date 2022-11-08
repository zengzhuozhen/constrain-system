package core

type constrain struct {
	informant   string
	newValue    func(source string)
	forgetValue func(source string)
}

func MakeTernaryConstraint(left, right []*Connector, forwardFunc, reverseFunc func(left, right []float64) float64) *constrain {
	newValue := func(source string) {
		var expectVariable *Connector
		var calFunc func(left, right []float64) float64
		leftValues := make([]float64, 0)
		rightValues := make([]float64, 0)
		for i, connector := range append(left, right...) {
			if connector.hasValue() { // connector has value ,needn't calculate
				if i < len(left) {
					leftValues = append(leftValues, *connector.value)
				} else {
					rightValues = append(rightValues, *connector.value)
				}
				continue
			}
			if expectVariable == nil { // connector hasn't valued, find out first expectVariable
				expectVariable = connector
				if i < len(left) {
					calFunc = forwardFunc
				} else {
					calFunc = reverseFunc
				}
				continue
			} else { //  connector hasn't valued ,now exists two expectVariable,can't calculate
				return
			}
		}
		if expectVariable == nil { // not exists unKnow expectVariable return
			return
		}
		// only one expectVariable,can do calculate
		expectVariable.SetValue(source, calFunc(leftValues, rightValues))
	}
	forgetValue := func(source string) {
		for _, connector := range append(left, right...) {
			connector.forgotValue(source)
		}
	}
	constraint := &constrain{
		informant:   left[0].informant,
		newValue:    newValue,
		forgetValue: forgetValue,
	}
	for _, connector := range append(left, right...) {
		connector.connect(constraint)
	}
	return constraint
}

func informAllExcept(source string, method string, constraints []*constrain) {
	for _, c := range constraints {
		if source != c.informant {
			switch method {
			case "new_value":
				c.newValue(source)
			case "forget_value":
				c.forgetValue(source)
			}
		}
	}
}
