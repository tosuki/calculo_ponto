package core

import "github.com/4mti/ponto/domain/model"

type ConfigLoader interface {
	LoadConfig(pathName string) (*model.Config, error)
	LoadConfigOrElse(pathName string, orElse *model.Config) *model.Config
	SaveConfig(config *model.Config, pathName string) error
}
