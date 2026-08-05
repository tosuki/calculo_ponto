package drawer

import (
	"bytes"
	"image/color"
	"log"

	"github.com/4mti/ponto/domain/core"
	"github.com/4mti/ponto/domain/model"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	text "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type OverlayDrawerImpl struct {
	cfg core.ConfigManager

	fontSource *text.GoTextFaceSource
	fontFace   *text.GoTextFace
}

func (this *OverlayDrawerImpl) DrawOverlayClock(
	str string,
	position model.OverlayPosition,
	color model.OverlayColor,
) error {
	return nil
}

func (this *OverlayDrawerImpl) RemoveOverlayClock() error {
	return nil
}

func (this *OverlayDrawerImpl) ResolveWindowPosition(position model.OverlayPosition) (int, int) {
	ww, wh := this.cfg.GetWindowSize()
	sw, sh := ebiten.Monitor().Size()

	switch position {
	case model.OverlayPositionTopLeft:
		return 0, 0
	case model.OverlayPositionBottomLeft:
		return 0, sh - wh
	case model.OverlayPositionTopRight:
		return sw - ww, 0
	case model.OverlayPositionBottomRight:
		return sw - ww, sh - wh
	default:
		log.Fatalf("Falha ao resolver a posição da janela. Posição %d é inválida", position)
	}

	return 0, 0
}

func (this *OverlayDrawerImpl) Update() error {
	return nil
}

func (this *OverlayDrawerImpl) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return this.cfg.GetWindowSize()
}

func (this *OverlayDrawerImpl) DrawText(screen *ebiten.Image, txt string) {
	textOptions := &text.DrawOptions{}

	textOptions.GeoM.Translate(20, 20)
	text.Draw(screen, "Hello world", this.fontFace, textOptions)
}

func (this *OverlayDrawerImpl) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0})
}

func (this *OverlayDrawerImpl) InitializeTextTools() error {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))

	if err != nil {
		return err
	}

	this.fontSource = fontSource
	this.fontFace = &text.GoTextFace{
		Source: this.fontSource,
		Size:   16,
	}

	return nil
}

func (this *OverlayDrawerImpl) Start() error {
	ww, wh := this.cfg.GetWindowSize()

	ebiten.SetWindowSize(ww, wh)
	ebiten.SetWindowPosition(this.ResolveWindowPosition(
		this.cfg.GetOverlayPosition(),
	))
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetScreenClearedEveryFrame(false)

	if err := this.InitializeTextTools(); err != nil {
		return err
	}

	return ebiten.RunGameWithOptions(this, &ebiten.RunGameOptions{
		ScreenTransparent: false,
	})
}

func NewOverlayDrawer(cfg core.ConfigManager) core.Drawer {
	return &OverlayDrawerImpl{
		cfg: cfg,
	}
}
