package main

import (
	"bytes"
	"gocfilter/basic"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
)

// JPGImage 表示存储 JPG 图像数据的结构体
type JPGImage struct {
	Tiles map[int][]JPGTile
	Data  []byte
}

// JPGTile 表示在 JPG 图像中存储瓦片数据的结构体
type JPGTile struct {
	Data [][]int
}

// ImageProcessor 是处理不同图像类型的接口
type ImageProcessor interface {
	ReadImage(filename string) error
	SaveImage(filename string) error
}

// JPGImageProcessor 是 JPG 图像的 ImageProcessor 实现
type JPGImageProcessor struct {
	Image JPGImage
}

// ReadImage 从指定文件读取 JPG 图像
func (processor *JPGImageProcessor) ReadImage(filename string) error {
	// 构建完整文件路径
	fullPath := filepath.Join(basic.ImportDir, filename)

	// 打开文件
	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码 JPG 图像
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 获取图像数据
	data, err := getImageData(img)
	if err != nil {
		return err
	}

	// 将图像数据存储到 JPGImage 的 Data 字段中
	processor.Image.Data = data

	return nil
}

// getImageData 从图像中提取字节数据
func getImageData(img image.Image) ([]byte, error) {
	// 这里简化为直接将图像编码为 JPG 格式的字节数据
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// SaveImage 将 JPG 图像保存到指定文件
func (processor *JPGImageProcessor) SaveImage(filename string) error {

	fullPath := filepath.Join(basic.ExportDir, filename)

	err := ioutil.WriteFile(fullPath, processor.Image.Data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// 创建 JPGImageProcessor 的实例
	jpgProcessor := &JPGImageProcessor{}

	// 指定要读取的 JPG 图像文件名
	filename := "cat.jpg"

	// 读取 JPG 图像
	err := jpgProcessor.ReadImage(filename)
	if err != nil {
		println("读取图像错误:", err)
		return
	}
	// Split into Tiles

	// Convolution

	// 指定要保存的 JPG 图像文件名
	saveFilename := "output.jpg"

	// 保存 JPG 图像
	err = jpgProcessor.SaveImage(saveFilename)
	if err != nil {
		println("保存图像错误:", err)
	}
}
