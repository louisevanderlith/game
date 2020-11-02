package core

import (
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type Experience struct {
	Type   ExperienceType
	Points int
	Date time.Time
}

func (o Experience) Valid() error {
	return validation.Struct(o)
}
