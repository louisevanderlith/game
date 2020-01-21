package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"

	"github.com/louisevanderlith/game/controllers"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.Unknown, mix.JSON, &controllers.Heroes{})

	//Hero
	/*
		heroCtrl := &controllers.Heroes{}
		heroGroup := routing.NewRouteGroup("hero", mix.JSON)
		heroGroup.AddRoute("Get Hero By ID", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, heroCtrl.View)
		heroGroup.AddRoute("Get All Heroes", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, heroCtrl.Get)

		e.AddBundle(heroGroup)

		//Level
		lvlCtrl := &controllers.Levels{}
		lvlGroup := routing.NewRouteGroup("level", mix.JSON)
		lvlGroup.AddRoute("Get All Heroes", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, lvlCtrl.Get)
		e.AddBundle(lvlGroup)*/
}
