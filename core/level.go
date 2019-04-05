package core

import (
	"github.com/louisevanderlith/husk"
)

type Level struct {
	Rank     int
	Required int
}

func (o Level) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

//GetRank finds the ranks you're not, and subtracts one
func GetRank(totalXP int) int {
	record, err := ctx.Levels.FindFirst(byNotMet(totalXP))

	if err != nil {
		return 0
	}

	lvl := record.Data().(*Level)

	return lvl.Rank - 1
}

//GetRequired returns the amount of XP required for the next level.
func GetRequired(totalXP int) int {
	record, err := ctx.Levels.FindFirst(byNotMet(totalXP))

	if err != nil {
		return 50
	}

	lvl := record.Data().(*Level)

	return lvl.Required - totalXP
}
