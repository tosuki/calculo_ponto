package render

import (
	"fmt"

	"github.com/4mti/ponto/src/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App struct {
	font        *NumberFont
	timer       *core.Timer
	pausedColor *NumberColor

	isDecorationEnabled       bool
	isMousePassthroughEnabled bool

	tick int
}

func (app *App) Update() error {
	app.tick++

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		fmt.Println("Tecla espaço pressionada.")

		if app.timer.IsPaused {
			app.timer.Resume()
		} else {
			app.timer.Pause()
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF6) {
		app.isDecorationEnabled = !app.isDecorationEnabled
		fmt.Printf("F6 Pressionado, mouse passthrough alterado pra %t\n", app.isMousePassthroughEnabled)
		ebiten.SetWindowDecorated(app.isDecorationEnabled)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF7) {
		app.isMousePassthroughEnabled = !app.isMousePassthroughEnabled

		fmt.Printf("F7 pressionado, mouse passthrough alterado pra %t\n", app.isMousePassthroughEnabled)
		ebiten.SetWindowMousePassthrough(app.isMousePassthroughEnabled)
	}

	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	row := (app.tick / 15) % 3
	text := app.timer.GetOutput()

	if app.timer.IsPaused {
		app.font.DrawText(text, screen, 0, 0, row, app.pausedColor)
	} else {
		app.font.DrawText(text, screen, 0, 0, row, nil)
	}
}

func (app *App) Layout(ow, oh int) (int, int) {
	tw, th := app.font.Measure(app.timer.GetOutput())

	//text width/height + a padding of 10 on each side
	return tw, th
}

func RunApp(timer *core.Timer, config *core.Config) error {
	font := NewNumberFont("assets/digits.png")

	app := &App{
		font:  font,
		timer: timer,

		pausedColor: &NumberColor{
			red:   100,
			green: 0,
			blue:  0,
		},

		isDecorationEnabled:       true,
		isMousePassthroughEnabled: false,
	}

	_, th := font.Measure(timer.GetOutput())
	ebiten.SetWindowSize(600, th)
	ebiten.SetWindowDecorated(true) // Sem bordas de janela
	ebiten.SetWindowFloating(true)  // ALWAYS ON TOP!

	if err := ebiten.RunGameWithOptions(app, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		return err
	}

	return nil
}
