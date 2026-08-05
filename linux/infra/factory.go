package infra

import (
	"github.com/4mti/ponto/domain/core"
	"github.com/4mti/ponto/infra/drawer"
	"github.com/4mti/ponto/infra/event"
	"github.com/4mti/ponto/infra/loader"
)

func NewOverlayDrawer(cfg core.ConfigManager) core.Drawer {
	return drawer.NewOverlayDrawer(cfg)
}

func NewEventListener() core.EventListener {
	return event.NewEventListener()
}

func NewConfigLoader() core.ConfigLoader {
	return &loader.ConfigLoaderImpl{}
}
