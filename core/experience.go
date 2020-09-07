package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type Experience struct {
	Type   ExperienceType
	Points int
}

func (o Experience) Valid() error {
	return validation.Struct(o)
}
