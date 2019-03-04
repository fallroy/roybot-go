package controller

import (
	"roybot/config"
	"time"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	t := time.Now()
	c.JSON(200, &versionResp{
		ReleaseTime:    config.Conf.Release.Time,
		ReleaseVersion: config.Conf.Release.Version,
		SystemTime:     t.Format("2006-01-02 15:04:05"),
	},
	)
}

type versionResp struct {
	ReleaseVersion string `json:"releaseVersion"`
	ReleaseTime    string `json:"releaseTime"`
	SystemTime     string `json:"systemTime"`
}
