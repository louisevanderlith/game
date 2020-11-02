package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"os"
	"reflect"
)

type context struct {
	Heroes husk.Table
	Levels husk.Table
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Heroes: husk.NewTable(Hero{}),
		Levels: husk.NewTable(Level{}),
	}
}

func Shutdown() {
	ctx.Heroes.Save()
}

func seed() {
	levels, err := levelSeeds()

	if err != nil {
		panic(err)
	}

	err = ctx.Levels.Seed(levels)

	if err != nil {
		panic(err)
	}
}

func levelSeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/levels.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Level
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}
