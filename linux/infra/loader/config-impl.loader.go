package loader

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/4mti/ponto/domain/model"
)

type ConfigLoaderImpl struct {
}

func (this *ConfigLoaderImpl) LoadConfig(pathName string) (*model.Config, error) {
	xdgConfigDir, err := os.UserConfigDir()

	if err != nil {
		log.Fatalf("Não foi possivel realizar a resolução do caminho da pasta de configuração do sistema: %s", err)
	}

	configDir := filepath.Join(xdgConfigDir, "ponto")
	err = os.MkdirAll(configDir, 0700)

	if err != nil {
		log.Fatalf("Não foi carregar a pasta de configuração: %s", err)
	}

	configPath := filepath.Join(configDir, pathName)
	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}

		log.Fatalf("Não foi possivel abrir o arquivo de configuração %s devido a: %s", configPath, err)
	}

	defer file.Close()

	var bytes []byte

	n, err := file.Read(bytes)

	if err != nil {
		log.Fatalf("Não foi possivel ler o arquivo devido a: %s", err)
	}

	if n <= 0 {
		return nil, nil
	}

	var config model.Config

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Não foi possivel realizar o decode do arquivo de configuração: %s", err)
	}

	return &config, nil
}

func (this *ConfigLoaderImpl) LoadConfigOrElse(pathName string, orElse *model.Config) *model.Config {
	config, err := this.LoadConfig(pathName)

	if err != nil {
		fmt.Println(err.Error())
		return orElse
	}

	if config == nil {
		return orElse
	}

	return config
}

func (this *ConfigLoaderImpl) SaveConfig(config *model.Config) error {
	return nil
}
