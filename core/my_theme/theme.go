package my_theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if variant == theme.VariantDark {
		switch name {
		case theme.ColorNameBackground:
			return color.RGBA{R: 50, G: 57, B: 73}
		case theme.ColorNameScrollBar:
			return color.RGBA{R: 31, G: 39, B: 60}
		case theme.ColorNameInputBackground:
			return color.RGBA{R: 31, G: 39, B: 60, A: 255}
		case theme.ColorNameButton:
			return color.RGBA{R: 39, G: 32, B: 62}
		case theme.ColorNameShadow:
			return color.RGBA{R: 31, G: 39, B: 60}
		case theme.ColorNamePlaceHolder:
			return color.RGBA{R: 50, G: 57, B: 73}
		}

	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	if name == theme.IconNameHome {
		fyne.NewStaticResource("myHome", []byte{})
	}

	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
