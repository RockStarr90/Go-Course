package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var width int = 640
var height int = 640
var cellNum int = 8

func main() {

	rectColor := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	draw.Draw(img, img.Bounds(), &image.Uniform{rectColor}, image.ZP, draw.Src)

	for x := 0; x < width; x++ {
		for y := height / cellNum; y < height; y += height / cellNum {
			img.Set(x, y, red)
		}
	}
	for y := 0; y < height; y++ {
		for x := width / cellNum; x < width; x += width / cellNum {
			img.Set(x, y, blue)
		}
	}
	png.Encode(file, img)
}
