package converter

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunshineplan/imgconv"
)

func TestConvert(t *testing.T) {
	originalExt := filepath.Ext("./sample/Go-Logo.jpg")

	ConvertImage("./sample/Go-Logo.jpg", "./sample/", "Go-Logo", imgconv.PNG)

	converterExt := filepath.Ext("./sample/Go-Logo.png")

	assert.NotNil(t, converterExt)
	assert.NotEqual(t, originalExt, converterExt)
}
