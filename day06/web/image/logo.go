package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	size      = 300
	imageName = "amazing_logo.png"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	cyan := color.RGBA{100, 200, 200, 0xff}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			switch {
			case x < size/2 && y < size/2:
				img.Set(x, y, cyan)
			case x >= size/2 && y >= size/2:
				img.Set(x, y, color.White)
			default:
			}
		}
	}

	f, err := os.Create(imageName)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
