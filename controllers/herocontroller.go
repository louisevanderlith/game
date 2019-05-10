package controllers

import "github.com/louisevanderlith/mango/control"

type HeroController struct {
	control.APIController
}

func NewHeroCtrl(ctrlMap *control.ControllerMap) *HeroController {
	result := &HeroController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

