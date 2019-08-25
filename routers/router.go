package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"

	"github.com/louisevanderlith/game/controllers"
)

func Setup(poxy resins.Epoxi) {
	//Hero
	heroCtrl := &controllers.Heroes{}
	heroGroup := routing.NewRouteGroup("hero", mix.JSON)
	heroGroup.AddRoute("Get Hero By ID", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, heroCtrl.View)
	heroGroup.AddRoute("Get All Heroes", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, heroCtrl.Get)

	poxy.AddGroup(heroGroup)

	//Level
	lvlCtrl := &controllers.Levels{}
	lvlGroup := routing.NewRouteGroup("level", mix.JSON)
	lvlGroup.AddRoute("Get All Heroes", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, lvlCtrl.Get)
	poxy.AddGroup(lvlGroup)

	/*ctrlmap := EnableFilters(s, host)

	heroCtrl := controllers.NewHeroCtrl(ctrlmap)
	beego.Router("/v1/hero/", heroCtrl)

	levelCtrl := controllers.NewLevelCtrl(ctrlmap)
	beego.Router("/v1/level", levelCtrl)*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner
	emptyMap["PUT"] = roletype.Owner

	ctrlmap.Add("/v1/hero", emptyMap)
	ctrlmap.Add("/v1/level", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
*/
