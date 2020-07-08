package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/game/core"
	"github.com/louisevanderlith/husk"
)

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	results, err := core.GetHeroes(1, 10)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

func SearchHeroes(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()

	results, err := core.GetHeroes(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

func ViewHeroes(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	record, err := core.GetHero(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(record.Data()))

	if err != nil {
		log.Println(err)
	}
}
