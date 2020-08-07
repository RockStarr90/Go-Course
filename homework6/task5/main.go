package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var width int = 640  // ширина шахматной доски (pixels)
var height int = 640 // высота шахматной доски (pixels)
var cellNum int = 8  // количество клеток по осям Х и Y
var deltaX int = 0   // начало следующей клетки по оси Х
var deltaY int = 0   // начало следующей клетки по оси Y
var countY int = 0   // счетчик рядов по оси Y

func main() {
	rectColor := color.RGBA{192, 192, 192, 255}
	black := color.RGBA{0, 0, 0, 255}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	file, err := os.Create("Chess board.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	draw.Draw(img, img.Bounds(), &image.Uniform{rectColor}, image.ZP, draw.Src)

loop:
	// отрисовываем черную клетку
	for x := 0 + deltaX; x < deltaX+width/cellNum; x++ {
		for y := 0 + deltaY; y < deltaY+height/cellNum; y++ {
			img.Set(x, y, black)
		}
	}
	// настраиваем начало отрисовки следующей клетки
	deltaX += 2 * width / cellNum
	if deltaX < width {
		goto loop
	} else {
		deltaX = 0
		countY += 1
		if countY == 8 {
			goto exit
		}
		deltaY += height / cellNum
		if countY%2 == 0 {
			deltaX = 0
		} else {
			deltaX = width / cellNum
		}
		goto loop
	}
exit:
	png.Encode(file, img)
}
