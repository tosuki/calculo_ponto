package core

import (
	"fmt"
	"log"
	"os"

	"github.com/4mti/ponto/domain/err"
	"github.com/4mti/ponto/domain/model"
)

type ConfigManager interface {
	Initialize() error
	SetOverlayPosition(position model.OverlayPosition) error
	SetOverlayColor(color model.OverlayColor) error
	ToggleOverlay() error
	GetOverlayPosition() model.OverlayPosition
	GetOverlayColor() model.OverlayColor
	IsOverlayEnabled() bool

	SetJourneyHours(hours int) error
	SetStartTime(time int64) error
	GetStartTime() int64
	AddClockIn(time int64, description string) error
	AddClockOut(time int64, description string) error

	GetCurrentClock() *model.Clock

	GetClocks() []*model.Clock
}

type ConfigManagerImpl struct {
	configFileState *model.Config
	state           model.GlobalState
	configLoader    ConfigLoader
	eventListener   EventListener
}

func (this *ConfigManagerImpl) GetConfigDir() string {
	userConfigDir, err := os.UserConfigDir()

	if err != nil {
		log.Fatalf("Não foi possivel realizar a montagem do diretório de configuração devido a um erro desconhecido: %s", err)
	}

	return fmt.Sprintf("%s/ponto", userConfigDir)
}

func (this *ConfigManagerImpl) GetMainConfigFilePath() string {
	return fmt.Sprintf("%s/config.json", this.GetConfigDir())
}

func (this *ConfigManagerImpl) refresh() error {
	if err := this.configLoader.SaveConfig(this.configFileState); err != nil {
		return err
	}

	return this.eventListener.Emit("refresh", this.state)
}

func (this *ConfigManagerImpl) Initialize() error {
	mainConfigPath := this.GetMainConfigFilePath()
	this.configFileState = this.configLoader.LoadConfigOrElse(mainConfigPath, model.NewConfig(false, model.OverlayPositionTopLeft, model.OverlayColorBlue, 100, true))

	this.state = *model.NewGlobalState(
		0,
		this.configFileState.JourneyHours,
		[]*model.Clock{},
		this.configFileState,
	)

	fmt.Println("Estado global da aplicação inicializado!")
	return this.refresh()
}

func (this *ConfigManagerImpl) SetJourneyHours(hours int) error {
	this.configFileState.JourneyHours = hours
	this.state.JourneyHours = hours

	return this.refresh()
}

func (this *ConfigManagerImpl) SetStartTime(time int64) error {
	this.state.StartTime = time
	this.configFileState.StartTime = time

	return this.refresh()
}

func (this *ConfigManagerImpl) GetStartTime() int64 {
	return this.state.StartTime
}

func (this *ConfigManagerImpl) AddClockIn(time int64, description string) error {
	if this.state.CurrentClock != nil {
		return err.ErrCurrentClockBusy
	}

	this.state.CurrentClock = model.NewClock(time, description)

	return this.refresh()
}

func (this *ConfigManagerImpl) AddClockOut(time int64, description string) error {
	if this.state.CurrentClock == nil {
		return err.ErrNoClockIn
	}

	this.state.CurrentClock.End()

	this.state.Clocks = append(this.state.Clocks, this.state.CurrentClock)
	this.state.CurrentClock = nil

	return this.refresh()
}

func (this *ConfigManagerImpl) GetCurrentClock() *model.Clock {
	return this.state.CurrentClock
}

func (this *ConfigManagerImpl) SetOverlayPosition(position model.OverlayPosition) error {
	if this.configFileState.OverlayPosition == position && this.state.OverlayPosition == position {
		fmt.Printf("Overlay já se encontra na posição %d\n", position)
		return nil
	}

	if !model.ValidateOverlayPosition(position) {
		return err.ErrInvalidPropertyValue
	}

	this.configFileState.OverlayPosition = position
	this.state.OverlayPosition = position

	return this.refresh()
}

func (this *ConfigManagerImpl) GetOverlayPosition() model.OverlayPosition {
	return this.state.OverlayPosition
}

func (this *ConfigManagerImpl) ToggleOverlay() error {
	this.configFileState.OverlayEnabled = !this.configFileState.OverlayEnabled
	this.state.OverlayEnabled = this.configFileState.OverlayEnabled

	return this.refresh()
}

func (this *ConfigManagerImpl) IsOverlayEnabled() bool {
	return this.state.OverlayEnabled
}

func (this *ConfigManagerImpl) SetOverlayColor(color model.OverlayColor) error {
	if !model.ValidateOverlayColor(color) {
		return err.ErrInvalidPropertyValue
	}

	this.configFileState.OverlayColor = color
	this.state.OverlayColor = color

	return this.refresh()
}

func (this *ConfigManagerImpl) GetOverlayColor() model.OverlayColor {
	return this.state.OverlayColor
}
