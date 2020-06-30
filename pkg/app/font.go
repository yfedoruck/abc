package app

import (
	"github.com/golang/freetype/truetype"
	"github.com/yfedoruck/abc/pkg/resources/fonts"
	"golang.org/x/image/font"
	"log"
)

func FontFace() font.Face {
	tt, err := truetype.Parse(fonts.Telelower_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	return truetype.NewFace(tt, &truetype.Options{
		Size:    27,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
