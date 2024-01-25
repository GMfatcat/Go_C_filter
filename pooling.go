package main

/*
#cgo CFLAGS: -Wall
#include "maxpooling.c"
*/
import "C"

import (
	"fmt"
	"gocfilter/process"
	"image"
	"image/color"
	"time"
	"unsafe"
)

func main() {

	// load image and turn into grayscale
	grayImg := process.LoadImage("cat.jpg")                     // small image
	grayImgL := process.LoadImage("./large_jpg/resize_cat.jpg") // large image

	// max pooling in c
	start := time.Now()
	result := MaxPoolingInC(grayImg, 3)
	duration := time.Since(start)
	fmt.Println("Max pooling process in c completed in:", duration.Nanoseconds())

	// max pooling in go
	start2 := time.Now()
	result2 := process.MaxPoolingInGo(grayImg, 3)
	duration2 := time.Since(start2)
	fmt.Println("Max pooling process in go completed in:", duration2.Nanoseconds())

	// Save Result
	process.SaveImage(result, "output_c_maxpooling.jpg")
	process.SaveImage(result2, "output_go_maxpooling.jpg")

	// 1000 TIME AVERAGE for small image
	cAvg := process.MeasureTime(func() {
		MaxPoolingInC(grayImg, 3)
	}, 1000)
	goAvg := process.MeasureTime(func() {
		process.MaxPoolingInGo(grayImg, 3)
	}, 1000)

	fmt.Printf("1000 times average for small image:\nC avg: %v\nGo avg: %v\n", cAvg, goAvg)

	// 1000 TIME AVERAGE for large image
	cAvg = process.MeasureTime(func() {
		MaxPoolingInC(grayImgL, 3)
	}, 1000)
	goAvg = process.MeasureTime(func() {
		process.MaxPoolingInGo(grayImgL, 3)
	}, 1000)

	fmt.Printf("1000 times average for large image:\nC avg: %v\nGo avg: %v\n", cAvg, goAvg)

	// End
	fmt.Println("Max pooling process completed.")
}

func MaxPoolingInC(img *image.Gray, poolSize int) *image.Gray {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	newWidth, newHeight := width/poolSize, height/poolSize
	newImg := image.NewGray(image.Rect(0, 0, newWidth, newHeight))

	// flatten image for C to max pooling
	input := make([]uint8, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			input[y*width+x] = img.GrayAt(x, y).Y
		}
	}

	// output image is also 1-D array
	output := make([]uint8, newWidth*newHeight)
	C.max_pooling((*C.uint8_t)(unsafe.Pointer(&input[0])), (*C.uint8_t)(unsafe.Pointer(&output[0])), C.int(width), C.int(height), C.int(poolSize))

	// transform back to 2D image
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			newImg.SetGray(x, y, color.Gray{Y: output[y*newWidth+x]})
		}
	}

	return newImg
}
