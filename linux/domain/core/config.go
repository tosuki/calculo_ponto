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
}

type ConfigManagerImpl struct {
	configFileState *model.Config
	state           model.GlobalState
	configLoader    ConfigLoader
	eventListener   EventListener
}

func (feat *ConfigManagerImpl) GetConfigDir() string {
	userConfigDir, err := os.UserConfigDir()

	if err != nil {
		log.Fatalf("Não foi possivel realizar a montagem do diretório de configuração devido a um erro desconhecido: %s", err)
	}

	return fmt.Sprintf("%s/ponto", userConfigDir)
}

func (feat *ConfigManagerImpl) GetMainConfigFilePath() string {
	return fmt.Sprintf("%s/config.json", feat.GetConfigDir())
}

func (feat *ConfigManagerImpl) refresh() error {
	if err := feat.configLoader.SaveConfig(feat.configFileState); err != nil {
		return err
	}

	return feat.eventListener.Emit("refresh", feat.state)
}

func (feat *ConfigManagerImpl) Initialize() error {
	mainConfigPath := feat.GetMainConfigFilePath()
	feat.configFileState = feat.configLoader.LoadConfigOrElse(mainConfigPath, model.NewConfig(false, model.OverlayPositionTopLeft, model.OverlayColorBlue, 100, true))

	feat.state = *model.NewGlobalState(
		0,
		0,
		[]*model.LunchPeriods{},
		feat.configFileState,
	)

	fmt.Println("Estado global da aplicação inicializado!")
	return feat.refresh()
}

func (feat *ConfigManagerImpl) SetOverlayPosition(position model.OverlayPosition) error {
	if feat.configFileState.OverlayPosition == position && feat.state.OverlayPosition == position {
		fmt.Printf("Overlay já se encontra na posição %d\n", position)
		return nil
	}

	if !model.ValidateOverlayPosition(position) {
		return err.ErrInvalidPropertyValue
	}

	feat.configFileState.OverlayPosition = position
	feat.state.OverlayPosition = position

	return feat.refresh()
}

func (feat *ConfigManagerImpl) GetOverlayPosition() model.OverlayPosition {
	return feat.state.OverlayPosition
}

func (feat *ConfigManagerImpl) ToggleOverlay() error {
	feat.configFileState.OverlayEnabled = !feat.configFileState.OverlayEnabled
	feat.state.OverlayEnabled = feat.configFileState.OverlayEnabled

	return feat.refresh()
}

func (feat *ConfigManagerImpl) IsOverlayEnabled() bool {
	return feat.state.OverlayEnabled
}

func (feat *ConfigManagerImpl) SetOverlayColor(color model.OverlayColor) error {
	if !model.ValidateOverlayColor(color) {
		return err.ErrInvalidPropertyValue
	}

	feat.configFileState.OverlayColor = color
	feat.state.OverlayColor = color

	return feat.refresh()
}

func (feat *ConfigManagerImpl) GetOverlayColor() model.OverlayColor {
	return feat.state.OverlayColor
}
