package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	_ "gogo/db"
	"gogo/log"
	"gogo/router"
)

func main() {
	port := flag.String("port", "8080", "server start port")

	flag.Parse()

	log.Info("port: ", *port)

	r := gin.New()

	// use middleware
	r.Use(gin.Logger(), gin.Recovery())

	// registered router
	router.Router(r)

	if err := r.Run(":" + *port); err != nil {
		log.Error(err)
	} else {
		log.Info("start server success")
	}
}
