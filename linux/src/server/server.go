package server

import (
	"github.com/4mti/ponto/src/core"
	"github.com/4mti/ponto/src/server/router"
	"github.com/gin-gonic/gin"
)

func StartServer(timer *core.Timer, config *core.Config) error {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := router.RegisterTimerRoutes(timer, config, r.Group("/timer")); err != nil {
		return err
	}

	if err := router.RegisterOverlayRoutes(config, r.Group("/overlay")); err != nil {
		return err
	}

	r.GET("/monitor", func(ctx *gin.Context) {
		monitor := config.GetMonitorDimensions()

		ctx.JSON(200, gin.H{
			"name":   monitor.Name,
			"height": monitor.Height,
			"width":  monitor.Width,
		})
	})
	r.Run()

	return nil
}
