package core

import (
	"errors"
	"fmt"
	"log"
)

var NoValueErr = errors.New("connector hadn't value")
var PerDefineUserSource = "user"

type Connector struct {
	Name        string
	value       *float64
	informant   string
	constraints []*constrain
	GetValue    func() float64
	SetValue    func(source string, value float64)
	ForgotValue func(source string)
	hasValue    func() bool
	connect     func(*constrain)
}

func MakeConnector(name string) *Connector {
	c := &Connector{Name: name}
	c.GetValue = func() float64 {
		if c.value == nil {
			panic(NoValueErr)
		}
		return *c.value
	}
	c.SetValue = func(source string, value float64) {
		if c.value != nil {
			if *c.value != value {
				panic(fmt.Sprintf("Contradiction detected:%.2f vs %.2f", *c.value, value))
			}
		} else {
			c.informant, c.value = source, &value
			informAllExcept(source, "new_value", c.constraints)
		}
	}
	c.ForgotValue = func(source string) {
		if c.informant == source {
			c.informant, c.value = "", nil
			if c.Name != "" {
				log.Default().Printf("%s is forgotten", c.Name)
			}
			informAllExcept(source, "forget_value", c.constraints)
		}
	}
	c.hasValue = func() bool {
		return c.value != nil
	}
	c.connect = func(constrain *constrain) {
		c.constraints = append(c.constraints, constrain)
	}
	return c
}
