package src

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type NumberFont struct {
	sheet *ebiten.Image
}

type NumberColor struct {
	red   float32
	green float32
	blue  float32
}

/**
* cw -> char width
* ch -> char height
 */
var (
	cw int
	ch int
)

func (nf *NumberFont) getCharFromImage(char rune, row int) *ebiten.Image {
	if row < 0 || row >= 3 {
		log.Fatalf("Tentativa de acessar uma fileira que não existe na imagem: %d\n", row)
		return nil
	}

	/**
	* 0 -> 0
	* 1 -> 1
	* Check the digits.png for more info
	 */
	var col int

	if char >= '0' && char <= '9' {
		col = int(char - '0')
	} else if char == ':' {
		col = 10
	} else {
		log.Fatalf("Tentativa de obter um caracter que não é suportado: %s\n", string(char))
		return nil
	}

	sx := col * cw
	sy := row * ch

	rect := image.Rect(sx, sy, sx+cw, sy+ch)

	return nf.sheet.SubImage(rect).(*ebiten.Image)
}

/**
* Supported digits [1-9], :
**/
func (nf *NumberFont) DrawText(
	text string,
	screen *ebiten.Image,
	x, y float64,
	row int,
	color *NumberColor,
) error {
	for i, char := range text {
		img := nf.getCharFromImage(char, row)

		if img == nil {
			continue
		}

		op := &ebiten.DrawImageOptions{}

		if color != nil {
			op.ColorScale.Scale(color.red, color.green, color.blue, 1.0)
		}

		op.GeoM.Translate(x+float64(i*cw), y)

		screen.DrawImage(img, op)
	}

	return nil
}

func (nf *NumberFont) Measure(text string) (tw, th int) {
	strLen := len(text)

	tw = cw * strLen
	th = ch

	return
}

func NewNumberFont(filepath string) *NumberFont {
	sheet, _, err := ebitenutil.NewImageFromFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	bounds := sheet.Bounds()
	cw = bounds.Dx() / 11
	ch = bounds.Dy() / 3

	return &NumberFont{
		sheet: sheet,
	}
}
