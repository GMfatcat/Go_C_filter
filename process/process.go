package process

import (
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "os"
    "time"
)

// max pooling in go
func MaxPoolingInGo(img *image.Gray, poolSize int) *image.Gray {
    bounds := img.Bounds()
    width, height := bounds.Dx(), bounds.Dy()
    newWidth, newHeight := width/poolSize, height/poolSize
    newImg := image.NewGray(image.Rect(0, 0, newWidth, newHeight))

    for y := 0; y < newHeight; y++ {
        for x := 0; x < newWidth; x++ {
            var max uint8
            for i := 0; i < poolSize; i++ {
                for j := 0; j < poolSize; j++ {
                    val := img.GrayAt(x*poolSize+i, y*poolSize+j).Y
                    if val > max {
                        max = val
                    }
                }
            }
            newImg.SetGray(x, y, color.Gray{Y: max})
        }
    }

    return newImg
}

func LoadImage(filename string) *image.Gray {

    var grayImg *image.Gray

    // 讀取影像
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error:", err)
        return grayImg
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("Error:", err)
        return grayImg
    }

    grayImg = image.NewGray(img.Bounds())

    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            grayImg.Set(x, y, img.At(x, y))
        }
    }

    return grayImg
}

func SaveImage(result *image.Gray, outputName string) {

    output, err := os.Create(outputName)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer output.Close()

    if err := jpeg.Encode(output, result, nil); err != nil {
        fmt.Println("Error:", err)
        return
    }

}

// use function wrapper to measure execution time
func MeasureTime(fn func(), iterations int) time.Duration {
    var totalDuration time.Duration
    for i := 0; i < iterations; i++ {
        start := time.Now()
        fn()
        duration := time.Since(start)
        totalDuration += duration
    }
    return totalDuration / time.Duration(iterations)
}
