package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
)

func main() {
	file, err := os.Open("brat.jpeg")
	if err != nil {
		fmt.Println("Error: brat.jpeg could not be opened")
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		return
	}

	bounds := img.Bounds()

	pixels := make([][]color.RGBA, bounds.Max.Y)

	for y := 0; y < bounds.Max.Y; y++ {
		pixels[y] = make([]color.RGBA, bounds.Max.X)
		for x := 0; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()
			pixels[y][x] = color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: 255,
			}
		}
	}

	s := ""
	reset := "\033[0m"
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[y]); x++ {
			pixel := pixels[y][x]
			color := fmt.Sprintf("\033[38;2;%d;%d;%dm", pixel.R, pixel.G, pixel.B)
			s += color + "██" + reset
		}
		s += "\n"
	}
	s += reset
	fmt.Println(s)
}