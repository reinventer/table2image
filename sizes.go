package table2image

import (
	"image"

	"golang.org/x/image/font"
)

type textSize struct {
	width, height int
}

func (t *Table) calculateSizes() [][]textSize {
	res := make([][]textSize, 0, len(t.data))
	for i, row := range t.data {
		res = append(res, make([]textSize, 0, len(row)))
		for _, text := range row {
			bounds, _ := font.BoundString(t.fontFace, text)
			res[i] = append(res[i], textSize{
				width:  (bounds.Max.X - bounds.Min.X).Floor(),
				height: (bounds.Max.Y - bounds.Min.Y).Floor(),
			})
		}
	}
	return res
}

type imageParameters struct {
	width, height int
	colWidths     []int
	rowHeights    []int
}

func (t *Table) getImageParameters(textSizes [][]textSize) imageParameters {
	if len(textSizes) == 0 {
		panic("empty table")
	}
	params := imageParameters{
		colWidths:  make([]int, len(textSizes[0]), len(textSizes[0])),
		rowHeights: make([]int, len(textSizes), len(textSizes)),
	}

	for i, row := range textSizes {
		var height int
		for j, sz := range row {
			if sz.width > params.colWidths[j] {
				params.colWidths[j] = sz.width
			}
			if sz.height > height {
				height = sz.height
			}
		}
		params.rowHeights[i] = height
	}

	params.width = sumInt(params.colWidths) + 2*len(params.colWidths)*t.paddingHorizontal
	params.height = sumInt(params.rowHeights) + 2*len(params.rowHeights)*t.paddingVertical
	return params
}

func (t *Table) findStartPoint(params imageParameters, textSizes [][]textSize, col, row int, horizontalAlign, verticalAlign Align) image.Point {
	startX := sumInt(params.colWidths[:col]) + (col*2+1)*t.paddingHorizontal
	startY := sumInt(params.rowHeights[:row+1]) + (row*2+1)*t.paddingVertical

	diffX := params.colWidths[col] - textSizes[row][col].width
	diffY := params.rowHeights[row] - textSizes[row][col].height

	switch horizontalAlign {
	case HorizontalAlignLeft:
	case HorizontalAlignCenter:
		startX += diffX / 2
	case HorizontalAlignRight:
		startX += diffX
	default:
		startX += diffX
	}

	switch verticalAlign {
	case VerticalAlignTop:
	case VerticalAlignMiddle:
		startY += diffY / 2
	case VerticalAlignBottom:
		startY += diffY
	default:
		startY += diffY
	}

	return image.Point{X: startX, Y: startY}
}

func sumInt(slice []int) (sum int) {
	for _, elem := range slice {
		sum += elem
	}
	return sum
}
