package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/game/core"
	"github.com/louisevanderlith/husk"
)

type Heroes struct {
	xontrols.APICtrl
}

func (req *Heroes) Get() {
	page, size := req.GetPageData()

	results := core.GetHeroes(page, size)

	req.Serve(http.StatusOK, nil, results)
}

func (req *Heroes) View() {
	key, err := husk.ParseKey(req.FindParam("uploadKey"))

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	record, err := core.GetHero(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, record)
}
