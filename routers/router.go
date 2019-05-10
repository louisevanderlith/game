// @APIVersion 1.0.0
// @Title Game API
// @Description API to control a users Levels and rewards
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/game/controllers"
	"github.com/louisevanderlith/mango"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilters(s)

	heroCtrl := controllers.NewHeroCtrl(ctrlmap)
	beego.Router("/v1/hero/", heroCtrl)

	levelCtrl := controllers.NewLevelCtrl(ctrlmap)
	beego.Router("/v1/level", levelCtrl)
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner
	emptyMap["PUT"] = roletype.Owner

	ctrlmap.Add("/v1/hero", emptyMap)
	ctrlmap.Add("/v1/level", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}), false)

	return ctrlmap
}
