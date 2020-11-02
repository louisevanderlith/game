package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type levelFilter func(obj Level) bool

func (f levelFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Level))
}

//returns the current level
func byNotMet(totalXP int) levelFilter {
	return func(obj Level) bool {
		return obj.Required > totalXP
	}
}
