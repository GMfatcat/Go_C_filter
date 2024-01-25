package basic

var ExportDir string = "./img"
var ImportDir string = "../example_img"

/* Basic Data Structure */

/*
Tile : small chunk of image
Tile number : record number to restore the pattern after processing
Ex:
1 2 3
4 5 6
7 8 9
*/
type JPGImage struct {
	Tiles map[int][]JPGTile
	Data  []byte
}

type JPGTile struct {
	Data [][]int
}

type BMPImage struct {
	Tiles map[int][]BMPTile
	Data  []byte
}

type BMPTile struct {
	Data [][]float64
}
