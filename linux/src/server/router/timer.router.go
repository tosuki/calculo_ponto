package router

import (
	"time"

	"github.com/4mti/ponto/src/core"
	"github.com/4mti/ponto/src/server/binding"
	"github.com/gin-gonic/gin"
)

func RegisterTimerRoutes(timer *core.Timer, config *core.Config, r *gin.RouterGroup) error {
	r.GET("/toggle", func(ctx *gin.Context) {
		if timer.IsPaused {
			timer.Resume()
		} else {
			timer.Pause()
		}

		ctx.Status(200)
	})

	r.PATCH("/config", func(ctx *gin.Context) {
		var body binding.DTOSetTimerConfig

		if err := ctx.ShouldBindBodyWithJSON(&body); err != nil || !body.Validate() {
			ctx.Status(400)
			return
		}

		if body.Journey != nil {
			timer.SetJourney(time.Duration(*body.Journey))
		}

		if body.StartedAt != nil {
			timer.SetJourney(time.Duration(*body.StartedAt))
		}

		ctx.JSON(201, gin.H{
			"journey":    timer.Journey,
			"started_at": timer.StartedAt,
			"output":     timer.GetOutput(),
		})
	})

	return nil
}
