package core

import "github.com/4mti/ponto/domain/model"

type Drawer interface {
	DrawOverlayClock(
		str string,
		position model.OverlayPosition,
		color model.OverlayColor,
	) error
	RemoveOverlayClock() error
	Start() error
}
