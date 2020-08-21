package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	//Hero
	/*
	ins := kong.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
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
	r := mux.NewRouter()

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "game.hero.search", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
