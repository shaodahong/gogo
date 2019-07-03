package main

import (
	"flag"
	"net/http"

	"gogo/ctx"
	_ "gogo/db"
	"gogo/log"
	"gogo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// parse command line args
	port := flag.String("port", "8080", "server start port")

	flag.Parse()

	log.Info("port: ", *port)

	r := gin.New()

	// use middleware
	r.Use(gin.Logger(), gin.Recovery())

	// handle no route request
	r.NoRoute(ctx.Handler(func(c *ctx.Ctx) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": "999999",
		})
	}))

	// handle no method request
	r.NoMethod(ctx.Handler(func(c *ctx.Ctx) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": "999999",
		})
	}))

	// register router
	router.Router(r)

	if err := r.Run(":" + *port); err != nil {
		log.Fatal(err)
	} else {
		log.Info("start server success")
	}
}
