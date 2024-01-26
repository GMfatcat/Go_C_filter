package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

const BlockSize = 15

func main() {

	img, err := readImage("../example_img/cat.jpg")
	if err != nil {
		fmt.Println("Error reading image:", err)
		return
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	for y := 0; y < height; y += BlockSize {
		for x := 0; x < width; x += BlockSize {

			blockBounds := image.Rect(x, y, x+BlockSize, y+BlockSize)

			blockImg := img.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(blockBounds)

			blockImg = applyConvolution(blockImg)

			for j := 0; j < BlockSize; j++ {
				for i := 0; i < BlockSize; i++ {
					draw.Draw(img.(draw.Image), blockBounds, blockImg, image.Point{}, draw.Over)
				}
			}
		}
	}

	err = saveImage("output.jpg", img)
	if err != nil {
		fmt.Println("Error saving image:", err)
	}
}

// applyConvolution mean kernel
func applyConvolution(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	for y := 1; y < bounds.Dy()-1; y++ {
		for x := 1; x < bounds.Dx()-1; x++ {
			// value in surrounding area
			c1 := img.At(x-1, y-1)
			c2 := img.At(x, y-1)
			c3 := img.At(x+1, y-1)
			c4 := img.At(x-1, y)
			c5 := img.At(x, y)
			c6 := img.At(x+1, y)
			c7 := img.At(x-1, y+1)
			c8 := img.At(x, y+1)
			c9 := img.At(x+1, y+1)

			r1, g1, b1, _ := c1.RGBA()
			r2, g2, b2, _ := c2.RGBA()
			r3, g3, b3, _ := c3.RGBA()
			r4, g4, b4, _ := c4.RGBA()
			r5, g5, b5, _ := c5.RGBA()
			r6, g6, b6, _ := c6.RGBA()
			r7, g7, b7, _ := c7.RGBA()
			r8, g8, b8, _ := c8.RGBA()
			r9, g9, b9, _ := c9.RGBA()

			avgR := (r1 + r2 + r3 + r4 + r5 + r6 + r7 + r8 + r9) / 9
			avgG := (g1 + g2 + g3 + g4 + g5 + g6 + g7 + g8 + g9) / 9
			avgB := (b1 + b2 + b3 + b4 + b5 + b6 + b7 + b8 + b9) / 9

			// new color
			newImg.Set(x, y, color.RGBA{
				uint8(avgR >> 8),
				uint8(avgG >> 8),
				uint8(avgB >> 8),
				255,
			})
		}
	}

	return newImg
}

func readImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	rgbaImg := image.NewRGBA(img.Bounds())
	draw.Draw(rgbaImg, rgbaImg.Bounds(), img, image.Point{}, draw.Src)

	return rgbaImg, nil
}

func saveImage(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}
