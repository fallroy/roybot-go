package main

import (
	"roybot/controller"

	"github.com/gin-gonic/gin"
)

func configRouter(r *gin.Engine) {
	r.POST("callback", controller.Callback)
	r.GET("callback", controller.TestCallback)
}
