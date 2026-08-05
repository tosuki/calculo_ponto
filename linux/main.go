package main

import (
	"fmt"

	"github.com/4mti/ponto/domain/core"
	"github.com/4mti/ponto/infra"
)

func main() {
	cfgLoader := infra.NewConfigLoader()
	eventListener := infra.NewEventListener()

	cfgManager := core.NewConfigManager(cfgLoader, eventListener)
	cfgManager.Initialize()

	overlay := infra.NewOverlayDrawer(cfgManager)

	if err := overlay.Start(); err != nil {
		fmt.Println(err.Error())
	}
}
