package table2image

import (
	"image"
	"image/color"
	"log"
	"os"
	"testing"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func TestExampleTable_WritePNG(t *testing.T) {
	ttFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("could not parse regular font: %s", err.Error())
	}

	face := truetype.NewFace(ttFont, &truetype.Options{Size: 14})

	table := NewTable(
		4,
		3,
		WithFontFace(face),
		WithTextColor(image.NewUniform(color.RGBA{R: 0xFF, G: 0x66, B: 0x66, A: 0xFF})),
		WithBackgroundColor(image.NewUniform(color.Gray{Y: 0xFF})),
		WithPaddings(5, 3),
		WithHorizontalAlign(HorizontalAlignLeft),
		WithVerticalAlign(VerticalAlignMiddle),
	)
	table.SetCell(0, 0, "Name")
	table.SetCell(1, 0, "Position")
	table.SetCell(2, 0, "Place")
	table.SetCell(3, 0, "House number")

	table.SetCell(0, 1, "John Dow")
	table.SetCell(1, 1, "Senior developer")
	table.SetCell(2, 1, "12345")
	table.SetCell(3, 1, "2")

	table.SetCell(0, 2, "Michael Johnson")
	table.SetCell(1, 2, "Actor")
	table.SetCell(2, 2, "54321")
	table.SetCell(3, 2, "5")

	f, err := os.OpenFile("example.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("could not create file: %s", err.Error())
	}
	defer f.Close()

	err = table.WritePNG(f)
	if err != nil {
		log.Fatalf("could not write png image: %s", err.Error())
	}
}
