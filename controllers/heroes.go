package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/game/core"
	"github.com/louisevanderlith/husk"
)

type Heroes struct {
}

func (req *Heroes) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetHeroes(1, 10)

	return http.StatusOK, results
}

func (req *Heroes) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()

	results := core.GetHeroes(page, size)

	return http.StatusOK, results
}

func (req *Heroes) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	record, err := core.GetHero(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, record
}
