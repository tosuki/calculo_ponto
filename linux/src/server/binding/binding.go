package binding

import (
	"github.com/4mti/ponto/src/core"
)

type DTOSetWindowPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}

func (d *DTOSetWindowPosition) Validate() bool {
	return d.X != nil && d.Y != nil
}

type DTOSetOverlayConfig struct {
	MousePassthrough *bool          `json:"mouse_passthrough"`
	WindowDecorated  *bool          `json:"window_decorated"`
	PausedColor      *core.RGBColor `json:"paused_color"`
}

func (d *DTOSetOverlayConfig) Validate() bool {
	return true
}

type DTOSetTimerConfig struct {
	StartedAt *int64 `json:"started_at"`
	Journey   *int64 `json:"journey"`
}

func (d *DTOSetTimerConfig) Validate() bool {
	return true
}
