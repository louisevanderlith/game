package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type Level struct {
	Rank     int
	Required int
}

func (o Level) Valid() error {
	return validation.Struct(o)
}

//GetRank finds the ranks you're not, and subtracts one
func GetRank(totalXP int) int {
	record, err := ctx.Levels.FindFirst(byNotMet(totalXP))

	if err != nil {
		return 0
	}

	lvl := record.GetValue().(*Level)

	return lvl.Rank - 1
}

//GetRequired returns the amount of XP required for the next level.
func GetRequired(totalXP int) int {
	record, err := ctx.Levels.FindFirst(byNotMet(totalXP))

	if err != nil {
		return 50
	}

	lvl := record.GetValue().(*Level)

	return lvl.Required - totalXP
}
