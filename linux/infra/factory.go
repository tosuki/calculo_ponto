package infra

import (
	"github.com/4mti/ponto/domain/core"
	"github.com/4mti/ponto/infra/loader"
)

func NewDrawer() core.Drawer {
	return nil
}

func NewEventListener() core.EventListener {
	return nil
}

func NewConfigLoader() core.ConfigLoader {
	return &loader.ConfigLoaderImpl{}
}
