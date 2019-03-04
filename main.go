package main

import (
	"roybot/config"
	"roybot/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()
	r.GET("version", controller.Version)
	configRouter(r)
	r.Run(config.Conf.ListenAddr)
}
