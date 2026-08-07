package server

import "github.com/4mti/ponto/src/render"

type DTOSetWindowPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (d *DTOSetWindowPosition) Validate() bool {
	return true
}

type DTOSetOverlayConfig struct {
	MousePassthrough bool               `json:"mouse_passthrough"`
	WindowDecorated  bool               `json:"window_decorated"`
	PausedColor      render.NumberColor `json:"paused_color"`
}

func (d *DTOSetOverlayConfig) Validate() bool {
	return true
}
