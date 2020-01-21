package hero

import (
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
	page, size := getPageData(c.Param("pagesize"))

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

func getPageData(pageData string) (int, int) {
	defaultPage := 1
	defaultSize := 10

	if len(pageData) < 2 {
		return defaultPage, defaultSize
	}

	pChar := []rune(pageData[:1])

	if len(pChar) != 1 {
		return defaultPage, defaultSize
	}

	page := int(pChar[0]) % 32
	pageSize, err := strconv.Atoi(pageData[1:])

	if err != nil {
		return defaultPage, defaultSize
	}

	return page, pageSize
}
