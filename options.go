package table2image

import (
	"image"

	"golang.org/x/image/font"
)

type OptionFunc func(*Table)

func WithFontFace(face font.Face) OptionFunc {
	return func(t *Table) {
		t.fontFace = face
	}
}

func WithTextColor(color *image.Uniform) OptionFunc {
	return func(t *Table) {
		t.textColor = color
	}
}

func WithBackgroundColor(color *image.Uniform) OptionFunc {
	return func(t *Table) {
		t.backgroundColor = color
	}
}

func WithPaddings(horizontalPadding, verticalPadding int) OptionFunc {
	return func(t *Table) {
		t.paddingHorizontal = horizontalPadding
		t.paddingVertical = verticalPadding
	}
}

func WithHorizontalAlign(align Align) OptionFunc {
	return func(t *Table) {
		t.horizontalAlign = align
	}
}

func WithVerticalAlign(align Align) OptionFunc {
	return func(t *Table) {
		t.verticalAlign = align
	}
}
