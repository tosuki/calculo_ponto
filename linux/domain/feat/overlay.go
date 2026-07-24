package feat

import (
	"fmt"

	"github.com/4mti/ponto/domain/core"
)

type OverlayFeat struct {
	eventListener core.EventListener
	config        core.ConfigManager
	drawer        core.Drawer
}

func (feat *OverlayFeat) Start() error {
	if !feat.config.IsOverlayEnabled() {
		fmt.Println("Overlay não se encontra ativa no momento. Portanto, não foi possivel iniciar o módulo.")
		return nil
	}

	go feat.refreshEventListener()

	for !feat.config.IsOverlayEnabled() {
		if err := feat.tick(); err != nil {
			return err
		}
	}

	return nil
}

func (feat *OverlayFeat) refreshEventListener() {
	refreshChan := feat.eventListener.On("refresh")

	for event := range refreshChan {
		fmt.Printf("Received signal to update the overlay (event name: %s)\n", event.Name)
		feat.tick()
	}
}

func (feat *OverlayFeat) tick() error {
	return nil
}

func NewOverlayFeat(
	eventListener core.EventListener,
	config core.ConfigManager,
	drawer core.Drawer,
) *OverlayFeat {
	return &OverlayFeat{
		eventListener: eventListener,
		config:        config,
		drawer:        drawer,
	}
}
