package router

import (
	"github.com/4mti/ponto/src/core"
	"github.com/gin-gonic/gin"
)

func RegisterTimerRoutes(timer *core.Timer, config *core.Config, r *gin.RouterGroup) error {
	r.GET("/toggle", func(ctx *gin.Context) {
		if timer.IsPaused {
			timer.Resume()
		} else {
			timer.Pause()
		}
	})

	r.PUT("/journey", func(ctx *gin.Context) {})

	return nil
}
