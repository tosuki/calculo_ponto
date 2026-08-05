package model

const (
	OverlayColorRed   = 0
	OverlayColorWhite = 1
	OverlayColorBlue  = 2
)

const (
	OverlayPositionTopLeft      = 0
	OverlayPositionBottomLeft   = 1
	OverlayPositionTopRight     = 2
	OverlayPositionBottomRight  = 3
	OverlayPositionBottomMiddle = 4
	OverlayPositionTopMiddle    = 5
)

type OverlayPosition = int
type OverlayColor = int

type Config struct {
	OverlayEnabled    bool  `json:"overlay_enabled"`
	OverlayPosition   int   `json:"overlay_position"`
	OverlayColor      int   `json:"overlay_color"`
	IsBorderEnabled   bool  `json:"is_border_enabled"`
	BackgroundOpacity int   `json:"background_opacity"`
	JourneyHours      int   `json:"journey_hours"`
	StartTime         int64 `json:"start_time"`
}

func (config *Config) Validate() bool {
	return ValidateOverlayBackgroundOpacity(config.BackgroundOpacity) &&
		ValidateOverlayPosition(config.OverlayPosition) &&
		ValidateOverlayColor(config.OverlayColor)
}

func ValidateOverlayBackgroundOpacity(opacity int) bool {
	return opacity >= 0 && opacity <= 100
}

func ValidateOverlayColor(color OverlayColor) bool {
	return color >= 0 && color <= 2
}

func ValidateOverlayPosition(position OverlayPosition) bool {
	return position >= 0 && position <= 5
}

func NewConfig(
	overlayEnabled bool,
	overlayPosition,
	overlayColor,
	backgroundOpacity int,
	isBorderEnabled bool,
) *Config {
	return &Config{
		OverlayEnabled:    overlayEnabled,
		OverlayColor:      overlayColor,
		OverlayPosition:   overlayPosition,
		BackgroundOpacity: backgroundOpacity,
		IsBorderEnabled:   isBorderEnabled,
	}
}
