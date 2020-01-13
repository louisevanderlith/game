package level

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Levels struct {
}

func Get(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, nil)
}
