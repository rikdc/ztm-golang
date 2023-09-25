package util

import (
	"image"
	"image/color"
)

func GetImageColors(img image.Image) map[color.Color]struct{} {
	colors := make(map[color.Color]struct{})
	var empty struct{}

}
