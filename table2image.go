package table2image

import (
	"image"
	"image/png"
	"io"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

type Table struct {
	cols, rows        int
	data              [][]string
	fontFace          font.Face
	textColor         *image.Uniform
	backgroundColor   *image.Uniform
	paddingHorizontal int
	paddingVertical   int
	horizontalAlign   Align
	verticalAlign     Align
}

func NewTable(cols, rows int, opts ...OptionFunc) *Table {
	data := make([][]string, rows, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]string, cols, cols)
	}

	ttFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		// should not panic
		panic(err)
	}

	table := &Table{
		cols:              cols,
		rows:              rows,
		data:              data,
		fontFace:          truetype.NewFace(ttFont, nil),
		textColor:         image.Black,
		backgroundColor:   image.White,
		paddingHorizontal: 5,
		paddingVertical:   2,
		horizontalAlign:   HorizontalAlignRight,
		verticalAlign:     VerticalAlignBottom,
	}

	for _, o := range opts {
		o(table)
	}

	return table
}

func (t *Table) SetCell(col, row int, s string) {
	if col >= t.cols || row >= t.rows {
		panic("wrong cell index")
	}
	t.data[row][col] = s
}

func (t *Table) WritePNG(w io.Writer) error {
	textSizes := t.calculateSizes()
	imageParams := t.getImageParameters(textSizes)

	img := image.NewRGBA(image.Rect(0, 0, imageParams.width, imageParams.height))
	draw.Draw(img, img.Bounds(), t.backgroundColor, image.Point{}, draw.Src)

	for i := 0; i < t.rows; i++ {
		for j := 0; j < t.cols; j++ {
			t.drawText(
				img,
				t.findStartPoint(imageParams, textSizes, j, i, t.horizontalAlign, t.verticalAlign),
				t.data[i][j],
			)
		}
	}
	return png.Encode(w, img)
}

func (t *Table) drawText(img draw.Image, point image.Point, s string) {
	drawer := font.Drawer{
		Dst:  img,
		Src:  t.textColor,
		Face: t.fontFace,
		Dot:  fixed.P(point.X, point.Y),
	}
	drawer.DrawString(s)
}
