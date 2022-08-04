package resources

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type LarkTheme struct {
	FontSize float32
}

func NewLarkTheme() *LarkTheme {
	t := LarkTheme{}
	t.FontSize = 13
	return &t
}

var _ fyne.Theme = (*LarkTheme)(nil)

func (t LarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t LarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t LarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return resourceCWindowsFontsSegoeuibTtf
	} else if style.Symbol {
		return resourceCWindowsFontsSeguisymTtf
	}
	return resourceCWindowsFontsSegoeuiTtf
}

func (t LarkTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 2
	case theme.SizeNameText:
		return t.FontSize
	}
	return theme.DefaultTheme().Size(name)
}
