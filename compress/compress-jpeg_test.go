package compress

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	originalFile, _ := os.Stat("./sample/Go-Logo.jpg")
	originalSize := originalFile.Size()

	CompressJPG("./compress")

	compressedFile, _ := os.Stat("./sample/Go-Logo-compress.jpg")
	compressedSize := compressedFile.Size()

	assert.Less(t, compressedSize, originalSize)
}
