package ctx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ctx struct {
	*gin.Context
}

func Handler(c func(*Ctx)) func(*gin.Context) {
	return func(g *gin.Context) {
		ctx := &Ctx{g}
		c(ctx)
	}
}

func (c Ctx) Success() {
	c.JSON(http.StatusOK, gin.H{"code": "000000"})
}
