package render

import (
	"fmt"

	"github.com/4mti/ponto/src/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App struct {
	font  *NumberFont
	timer *core.Timer
	cfg   *core.Config
	tick  int
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
		app.cfg.SetWindowDecorated(!app.cfg.IsWindowDecorated())
		fmt.Printf("F6 Pressionado, mouse passthrough alterado pra %t\n", app.cfg.IsWindowDecorated())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF7) {
		app.cfg.SetMousePassthrough(!app.cfg.IsMousePassthroughEnabled())
		fmt.Printf("F7 pressionado, mouse passthrough alterado pra %t\n", app.cfg.IsMousePassthroughEnabled())
		ebiten.SetWindowMousePassthrough(app.cfg.IsMousePassthroughEnabled())
	}

	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	row := (app.tick / 15) % 3
	text := app.timer.GetOutput()

	if app.timer.IsPaused {
		app.font.DrawText(text, screen, 0, 0, row, app.cfg.GetPausedColor())
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
		cfg:   config,
	}

	_, th := font.Measure(timer.GetOutput())

	config.SetWindowDecorated(true)
	config.SetPausedColor(&core.RGBColor{
		Red:   100,
		Green: 0,
		Blue:  0,
	})
	ebiten.SetWindowSize(600, th)
	ebiten.SetWindowFloating(true)

	if err := ebiten.RunGameWithOptions(app, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		return err
	}

	return nil
}
