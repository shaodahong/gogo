package router

import (
	"github.com/gin-gonic/gin"
	"gogo/ctx"
)

func Router(r *gin.Engine) {
	r.GET("/ping", ctx.Handler(func(c *ctx.Ctx) {
		c.Success()
	}))
}
