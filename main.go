package main

import (
	"log"
	"os"

	"github.com/louisevanderlith/game/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/game/core"
)

func main() {
	mode := os.Getenv("RUNMODE")
	pubPath := os.Getenv("KEYPATH")

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
