package hero

import (
	"github.com/louisevanderlith/droxo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/game/core"
	"github.com/louisevanderlith/husk"
)

func Get(c *gin.Context) {
	results := core.GetHeroes(1, 10)

	c.JSON(http.StatusOK, results)
}

func Search(c *gin.Context) {
	page, size := droxo.GetPageData(c.Param("pagesize"))

	results := core.GetHeroes(page, size)

	c.JSON(http.StatusOK, results)
}

func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	record, err := core.GetHero(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, record)
}