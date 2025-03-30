package screen

import (
	"image"
	"log"

	scr "github.com/kbinani/screenshot"
)

const (
	x = 20   //start position
	y = 100  //start position
	w = 1200 // picture wigth
	h = 800  // picture height
)

func MekeScreen() *image.RGBA {
	im, err := scr.Capture(x, y, w, h)
	if err != nil {
		log.Fatal("problem with create screenshot")
	}
	return im
}
