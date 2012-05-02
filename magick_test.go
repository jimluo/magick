package magick

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
	"os"
	"testing"
)

func setupImage(t *testing.T) (image *MagickImage) {
	filename := "test/heart_original.png"
	image, error := NewFromFile(filename)
	assert.T(t, error == nil)
	assert.T(t, image != nil)
	return
}

func TestImageFromFile(t *testing.T) {
	filename := "test/heart_original.png"
	image, error := NewFromFile(filename)
	assert.T(t, error == nil)
	assert.T(t, image != nil)

	bad_filename := "test/heart_whatwhat.png"
	image, error = NewFromFile(bad_filename)
	assert.T(t, error != nil)
}

func TestImageFromBlob(t *testing.T) {
	filename := "test/heart_original.png"
	source, _ := ioutil.ReadFile(filename)
	image, error := NewFromBlob(source, "png")
	assert.T(t, error == nil)
	assert.T(t, image != nil)
}

func TestParseGeometry(t *testing.T) {
	image := setupImage(t)
	geometry, err := image.ParseGeometry("100x100>")
	assert.T(t, err == nil)
	assert.T(t, geometry != nil)
	assert.Equal(t, 100, geometry.Width)
}

func TestDestroy(t *testing.T) {
	image := setupImage(t)
	assert.T(t, image.Destroy() == nil)
	assert.T(t, image.Image == nil)
}

func TestResize(t *testing.T) {
	image := setupImage(t)
	err := image.Resize("100x100!")
	assert.T(t, err == nil)
	assert.Equal(t, 100, image.Width())
	assert.Equal(t, 100, image.Height())

	image = setupImage(t)
	err = image.Resize("blurgh")
	assert.T(t, err != nil)
}

func TestCrop(t *testing.T) {
	image := setupImage(t)
	err := image.Crop("100x100!+10+10")
	assert.T(t, err == nil)
	assert.Equal(t, 100, image.Width())
	assert.Equal(t, 100, image.Height())

	image = setupImage(t)
	err = image.Crop("blurgh")
	assert.T(t, err != nil)
}

func TestShadow(t *testing.T) {
	image := setupImage(t)
	err := image.Shadow("#000", 75, 2, 0, 0)
	assert.T(t, err == nil)
}

func TestFillBackgroundColor(t *testing.T) {
	image := setupImage(t)
	err := image.FillBackgroundColor("#CCC")
	assert.T(t, err == nil)
}

func TestToBlob(t *testing.T) {
	image := setupImage(t)
	bytes, err := image.ToBlob("png")
	assert.T(t, err == nil)
	assert.T(t, bytes != nil)
	assert.T(t, len(bytes) > 0)
}

func TestToFile(t *testing.T) {
	image := setupImage(t)
	filename := "test/test_out.png"
	err := os.Remove(filename)
	err = image.ToFile(filename)
	assert.T(t, err == nil)
	file, err := os.Open(filename)
	assert.T(t, err == nil)
	defer file.Close()
	stat, err := file.Stat()
	assert.T(t, stat != nil)
}
