package converter

import (
	"log"

	"github.com/sunshineplan/imgconv"
)

func ConvertImage(srcPath string, dstPath string, dstFileName string, targetFormat imgconv.Format) {
	srcImage, err := imgconv.Open(srcPath)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}

	dstModifiedName := dstPath + dstFileName + "." + targetFormat.String()
	err = imgconv.Save(dstModifiedName, srcImage, &imgconv.FormatOption{Format: targetFormat})
	if err != nil {
		log.Fatalf("Failed to write image: %v", err)
	}
}
