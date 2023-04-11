package compress

import (
	"fmt"
	"os"

	"image/jpeg"
)

func CompressJPG(currentDir string) {
	jpegOption := jpeg.Options{
		Quality: 75,
	}

	os.Chdir(currentDir)
	imageFile, err := os.Open("./sample/Go-Logo.jpg")
	if err != nil {
		panic(err.Error())
	}
	defer imageFile.Close()

	outputFile, _ := os.Create("./sample/Go-Logo-compress.jpg")

	imageInput, err := jpeg.Decode(imageFile)
	if err != nil {
		fmt.Printf("PNG image decoding failed, error = %s", err.Error())
		return
	}

	jpeg.Encode(outputFile, imageInput, &jpegOption)
	fmt.Println("Finished compressing JPEG image")
}
