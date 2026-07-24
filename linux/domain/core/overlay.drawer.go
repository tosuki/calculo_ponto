package core

import "github.com/4mti/ponto/domain/model"

type Drawer interface {
	DrawOverlayClock(
		time, timeLeft string,
		overtimeMode bool,
		mode int,
		position model.OverlayPosition,
	) error

	RemoveOverlayClock() error
}
