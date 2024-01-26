package main

import (
	"fmt"
	"gocfilter/process"
	"time"
)

func main() {

	// load image and turn into grayscale
	grayImg := process.LoadImage("cat.jpg")                     // small image
	grayImgL := process.LoadImage("./large_jpg/resize_cat.jpg") // large image

	// max pooling in c
	start := time.Now()
	result := process.MaxPoolingInC(grayImg, 3)
	duration := time.Since(start)
	fmt.Println("Max pooling process in c completed in:", duration.Nanoseconds())

	// max pooling in go
	start2 := time.Now()
	result2 := process.MaxPoolingInGo(grayImg, 3)
	duration2 := time.Since(start2)
	fmt.Println("Max pooling process in go completed in:", duration2.Nanoseconds())

	// max pooling in goroutine
	start3 := time.Now()
	result3 := process.MaxPoolingInGoRoutine(grayImg, 3)
	duration3 := time.Since(start3)
	fmt.Println("Max pooling process in goroutine completed in:", duration3.Nanoseconds())

	// Save Result
	process.SaveImage(result, "output_c_maxpooling.jpg")
	process.SaveImage(result2, "output_go_maxpooling.jpg")
	process.SaveImage(result3, "output_goroutine_maxpooling.jpg")

	// 1000 TIME AVERAGE for small image
	cAvg := process.MeasureTime(func() {
		process.MaxPoolingInC(grayImg, 3)
	}, 1000)
	goAvg := process.MeasureTime(func() {
		process.MaxPoolingInGo(grayImg, 3)
	}, 1000)
	goRAvg := process.MeasureTime(func() {
		process.MaxPoolingInGoRoutine(grayImg, 3)
	}, 1000)

	fmt.Printf("1000 times average for small image:\nC avg: %v\nGo avg: %v\nGoroutine avg: %v\n",
		cAvg, goAvg, goRAvg)

	// 1000 TIME AVERAGE for large image
	cAvg = process.MeasureTime(func() {
		process.MaxPoolingInC(grayImgL, 3)
	}, 1000)
	goAvg = process.MeasureTime(func() {
		process.MaxPoolingInGo(grayImgL, 3)
	}, 1000)
	goRAvg = process.MeasureTime(func() {
		process.MaxPoolingInGoRoutine(grayImgL, 3)
	}, 1000)

	fmt.Printf("1000 times average for large image:\nC avg: %v\nGo avg: %v\nGoroutine avg: %v\n",
		cAvg, goAvg, goRAvg)

	// End
	fmt.Println("Max pooling process completed.")
}
