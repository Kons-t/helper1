package save

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"
)

type Save struct {
	Name string
	Path string
}

const (
	outputDir = "/Users/konstantinpetrenko/Desktop" // place  where save screenshot

)

func NewSave() *Save {
	return &Save{fmt.Sprintf("screenshot_%s.png", time.Now().Format("20060102_150405")), outputDir}
}

func (s *Save) Save(img *image.RGBA) error {
	filePath := filepath.Join(s.Path, s.Name)
	file, err := os.Create(filePath)
	if err != nil {
		return (fmt.Errorf("problem with creating file: %v", err))
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return (fmt.Errorf("problem with encoding file: %v", err))
	}
	return nil
}
