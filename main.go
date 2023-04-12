package main

import (
	"github.com/jwiryo/go-squeeze-image/compress"
	"github.com/jwiryo/go-squeeze-image/converter"
	"github.com/sunshineplan/imgconv"
)

func main() {
	compress.CompressJPG("./compress")
	converter.ConvertImage("./compress/sample/Go-Logo.jpg", "./compress/sample/", "Go-Logo", imgconv.PNG)
}
