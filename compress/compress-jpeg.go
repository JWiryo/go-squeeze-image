package compress

import (
	"fmt"
	"os"

	"image/jpeg"
)

func CompressJPG() {
	jpegOption := jpeg.Options{
		Quality: 75,
	}

	imageFile, err := os.Open("./compress/sample/Go-Logo.jpg")
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}
	defer imageFile.Close()

	outputFile, _ := os.Create("./compress/sample/Go-Logo-compress.jpg")

	imageInput, err := jpeg.Decode(imageFile)
	if err != nil {
		fmt.Printf("PNG image decoding failed, error = %s", err.Error())
		return
	}

	jpeg.Encode(outputFile, imageInput, &jpegOption)
	fmt.Println("Finished compressing JPEG image")
}
