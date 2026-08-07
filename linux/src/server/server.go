package server

import (
	"github.com/4mti/ponto/src/core"
	"github.com/gin-gonic/gin"
)

func StartServer(timer *core.Timer, config *core.Config) error {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/toggle", func(ctx *gin.Context) {
		if timer.IsPaused {
			timer.Resume()
		} else {
			timer.Pause()
		}

		ctx.JSON(200, gin.H{
			"is_paused": timer.IsPaused,
		})
	})

	router.GET("/monitor", func(ctx *gin.Context) {
		monitor := config.GetMonitorDimensions()

		ctx.JSON(200, gin.H{
			"name":   monitor.Name,
			"height": monitor.Height,
			"width":  monitor.Width,
		})
	})

	router.POST("/monitor/position", func(ctx *gin.Context) {

	})

	router.Run()

	return nil
}
