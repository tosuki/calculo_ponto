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

	router.PUT("/overlay/config", func(ctx *gin.Context) {
		var body DTOSetOverlayConfig

		if err := ctx.ShouldBindBodyWithJSON(&body); err != nil || !body.Validate() {
			ctx.Status(400)
			return
		}

	})

	router.PUT("/overlay/position", func(ctx *gin.Context) {
		var body DTOSetWindowPosition

		if err := ctx.ShouldBindBodyWithJSON(&body); err != nil || !body.Validate() {
			ctx.Status(400)
			return
		}

		config.SetPosition(body.X, body.Y)
		wx, wy := config.GetPosition()

		ctx.JSON(200, gin.H{
			"x": wx,
			"y": wy,
		})
	})

	router.Run()

	return nil
}
