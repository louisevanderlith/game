package controllers

import (
	"github.com/louisevanderlith/mango/control"
)

type LevelController struct {
	control.APIController
}

func NewLevelCtrl(ctrlMap *control.ControllerMap) *LevelController {
	result := &LevelController{}
	result.SetInstanceMap(ctrlMap)

	return result
}
