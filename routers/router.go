package routers

import (
	"github.com/louisevanderlith/droxolite"

	"github.com/louisevanderlith/game/controllers"
)

func Setup(poxy *droxolite.Epoxy) {
	//Hero
	heroCtrl := &controllers.HeroController{}
	heroGroup := droxolite.NewRouteGroup("hero", heroCtrl)

	poxy.AddGroup(heroGroup)

	//Level
	lvlCtrl := &controllers.LevelController{}
	lvlGroup := droxolite.NewRouteGroup("level", lvlCtrl)

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
