package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Heroes husk.Tabler
	Levels husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Heroes: husk.NewTable(new(Hero)),
		Levels: husk.NewTable(new(Level)),
	}
}

func Shutdown() {
	ctx.Heroes.Save()
}

func seed() {
	//ctx.Heroes - Levels
	err := ctx.Levels.Seed("db/levels.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Levels.Save()
}
