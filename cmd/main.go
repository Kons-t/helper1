package main

import (
	"github.com/Kons-t/helper/internal/save"
	"github.com/Kons-t/helper/internal/screen"
)

func main() {
	// runtime.LockOSThread()

	// hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyZ)
	// err := hk.Register()
	// if err != nil {
	// 	log.Fatalf("hotkey: failed to register hotkey: %v", err)
	// }

	save := save.NewSave()
	// <-hk.Keydown()
	im := screen.MekeScreen()
	err := save.Save(im)
	if err != nil {
		panic(err)
	}

}
