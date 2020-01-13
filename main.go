package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/game/controllers/hero"
	"github.com/louisevanderlith/game/controllers/level"
	"github.com/louisevanderlith/game/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()

	r.GET("/hero/:key", hero.View)

	// heroes := r.Group("/hero")
	// heroes.POST("", hero.Create)
	// heroes.PUT("/:key", hero.Update)
	// heroes.DELETE("/:key", hero.Delete)

	r.GET("/heroes", hero.Get)
	r.GET("/heroes/:pagesize/*hash", hero.Search)

	// r.GET("/level/:key", level.View)

	// levels := r.Group("/level")
	// levels.POST("", level.Create)
	// levels.PUT("/:key", level.Update)
	// levels.DELETE("/:key", level.Delete)

	r.GET("/levels", level.Get)
	// r.GET("/levels/:pagesize/*hash", level.Search)

	err := r.Run(":8100")

	if err != nil {
		panic(err)
	}
}

// func main() {
// 	keyPath := os.Getenv("KEYPATH")
// 	pubName := os.Getenv("PUBLICKEY")
// 	host := os.Getenv("HOST")
// 	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
// 	appName := os.Getenv("APPNAME")
// 	pubPath := path.Join(keyPath, pubName)

// 	// Register with router
// 	srv := bodies.NewService(appName, "", pubPath, host, httpport, servicetype.API)

// 	routr, err := do.GetServiceURL("", "Router.API", false)

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = srv.Register(routr)

// 	if err != nil {
// 		panic(err)
// 	}

// 	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, "none"))
// 	routers.Setup(poxy)
// 	poxy.EnableCORS(host)

// 	core.CreateContext()
// 	defer core.Shutdown()

// 	err = droxolite.Boot(poxy)

// 	if err != nil {
// 		panic(err)
// 	}
// }
