package feat

import (
	"fmt"
	"time"

	"github.com/4mti/ponto/domain/core"
)

type OverlayFeat struct {
	eventListener core.EventListener
	config        core.ConfigManager
	drawer        core.Drawer
}

func (this *OverlayFeat) Start() error {
	if !this.config.IsOverlayEnabled() {
		fmt.Println("Overlay não se encontra ativa no momento. Portanto, não foi possivel iniciar o módulo.")
		return nil
	}

	go this.refreshEventListener()

	for this.config.IsOverlayEnabled() {
		this.eventListener.Emit("refresh", nil)
		time.Sleep(time.Second * 1)
	}

	return nil
}

func (this *OverlayFeat) refreshEventListener() {
	refreshChan := this.eventListener.On("refresh")

	for event := range refreshChan {
		fmt.Printf("Received signal of event %s\n", event.Name)
		this.tick()
	}
}

func (this *OverlayFeat) tick() error {
	if !this.config.IsOverlayEnabled() {
		return nil
	}

	currentClock := this.config.GetCurrentClock()

	if currentClock != nil {
		return this.DrawCurrentClockTimer()
	}

	startTime := this.config.GetStartTime()

	if startTime <= 0 {
		return this.DrawIdleClock()
	}

	return this.DrawWorkTime(startTime)
}

func (this *OverlayFeat) DrawIdleClock() error {
	return this.Draw("00:00:00")
}

func (this *OverlayFeat) Draw(str string) error {
	color := this.config.GetOverlayColor()
	position := this.config.GetOverlayPosition()

	return this.drawer.DrawOverlayClock(str, position, color)
}

func (this *OverlayFeat) DrawWorkTime(startTimeUnix int64) error {
	now := time.Now()
	startTime := time.Unix(startTimeUnix, 0)

	if err := this.Draw(
		this.FormatDurationToString(now.Sub(startTime)),
	); err != nil {
		return err
	}

	return nil
}

func (this *OverlayFeat) FormatDurationToString(duration time.Duration) string {
	return duration.Abs().String()
}

func (this *OverlayFeat) DrawCurrentClockTimer() error {
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
