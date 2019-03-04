package controller

import (
	"fmt"

	"roybot/config"

	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {
	fmt.Printf("Hi %+v", config.Conf)
}
