package ui

import (
	"fyne.io/fyne/v2"
	"ztm.io/pixl/apptype"
	"ztm.io/pixl/pxcanvas"
	"ztm.io/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
