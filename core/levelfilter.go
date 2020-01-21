package core

import "github.com/louisevanderlith/husk"

type levelFilter func(obj Level) bool

func (f levelFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Level))
}

//returns the current level
func byNotMet(totalXP int) levelFilter {
	return func(obj Level) bool {
		return obj.Required > totalXP
	}
}
