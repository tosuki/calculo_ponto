package loader

import "github.com/4mti/ponto/domain/model"

type ConfigLoaderImpl struct {
}

func (this *ConfigLoaderImpl) LoadConfig(pathName string) (*model.Config, error) {
	return nil, nil
}

func (this *ConfigLoaderImpl) LoadConfigOrElse(pathName string, orElse *model.Config) *model.Config {
	return orElse
}

func (this *ConfigLoaderImpl) SaveConfig(config *model.Config) error {
	return nil
}
