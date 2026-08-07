package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Config struct {
	px int
	py int

	pausedColor *RGBColor
}

type MonitorMetadata struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (c *Config) GetPosition() (int, int) {
	return c.px, c.py
}

func (c *Config) SetPosition(x, y int) {
	c.px = x
	c.py = y

	ebiten.SetWindowPosition(c.px, c.py)
}

func (c *Config) SetWindowDecorated(value bool) {
	ebiten.SetWindowDecorated(value)
}

func (c *Config) IsWindowDecorated() bool {
	return ebiten.IsWindowDecorated()
}

func (c *Config) SetMousePassthrough(value bool) {
	ebiten.SetWindowMousePassthrough(value)
}

func (c *Config) IsMousePassthroughEnabled() bool {
	return ebiten.IsWindowMousePassthrough()
}

func (c *Config) GetMonitorDimensions() MonitorMetadata {
	monitor := ebiten.Monitor()

	mw, mh := monitor.Size()
	name := monitor.Name()

	return MonitorMetadata{
		Width:  mw,
		Height: mh,
		Name:   name,
	}
}

func (c *Config) GetPausedColor() *RGBColor {
	return c.pausedColor
}

func (c *Config) SetPausedColor(color *RGBColor) {
	c.pausedColor = color
}

func NewConfig() *Config {
	return &Config{}
}
