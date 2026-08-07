package core

import "github.com/hajimehoshi/ebiten/v2"

type Config struct {
	Px int
	Py int
	Sw int
	Sy int
}

type MonitorMetadata struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (c *Config) GetPosition() (int, int) {
	return c.Px, c.Py
}

func (c *Config) SetPosition(x, y int) {
	c.Px = x
	c.Py = y
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

func NewConfig() *Config {
	return &Config{
		Px: 0,
		Py: 0,
	}
}
