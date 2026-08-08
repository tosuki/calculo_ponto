package router

import (
	"github.com/4mti/ponto/src/core"
	"github.com/4mti/ponto/src/server/binding"
	"github.com/gin-gonic/gin"
)

func RegisterOverlayRoutes(config *core.Config, r *gin.RouterGroup) error {
	r.PUT("/overlay/config", func(ctx *gin.Context) {
		var body binding.DTOSetOverlayConfig

		if err := ctx.ShouldBindBodyWithJSON(&body); err != nil || !body.Validate() {
			ctx.Status(400)
			return
		}

		if body.MousePassthrough != nil {
			config.SetMousePassthrough(*body.MousePassthrough)
		}

		if body.WindowDecorated != nil {
			config.SetWindowDecorated(*body.WindowDecorated)
		}

		if body.PausedColor != nil {
			config.SetPausedColor(body.PausedColor)
		}
	})

	r.PUT("/overlay/position", func(ctx *gin.Context) {
		var body binding.DTOSetWindowPosition

		if err := ctx.ShouldBindBodyWithJSON(&body); err != nil || !body.Validate() {
			ctx.Status(400)
			return
		}

		wx, wy := config.GetPosition()

		if body.X != nil {
			wx = *body.X
		}

		if body.Y != nil {
			wy = *body.Y
		}

		config.SetPosition(wx, wy)
		ctx.JSON(200, gin.H{
			"x": wx,
			"y": wy,
		})
	})

	return nil
}
