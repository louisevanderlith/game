package core

import "github.com/louisevanderlith/husk"

type Experience struct {
	Type   ExperienceType
	Points int
}

func (o Experience) Valid() error {
	return husk.ValidateStruct(&o)
}
