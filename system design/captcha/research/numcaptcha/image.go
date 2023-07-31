package numcaptcha

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
)

const (
	// StdWidth standard width of a captcha image
	StdWidth = 240
	// StdHeight standard width of a captcha image
	StdHeight = 80
	// Maximum absolute skew factor of a single digit
	maxSkew = 0.7
	// Number of background circles
	circleCount = 20
)

// Image captcha
type Image struct {
	*image.Paletted
	numWidth  int
	numHeight int
	dotSize   int
	rng       siprng
}

// NewImage returns a new captcha image of the given width and height with
// the given digits, where each digit must be in range 0-9
func NewImage(id string, digits []byte, width, height int) *Image {
	m := new(Image)

	return m
}

func (m *Image) getRandomPalette() color.Palette {
	p := make([]color.Color, circleCount+1)
	// Transparent color
	p[0] = color.RGBA{0xFF, 0xFF, 0xFF, 0x00}
	// Primary color
	// prim := color.RGBA{
	// 	unit8(m.rng.Intn(129)),
	// 	unit8(m.rng.Intn(129)),
	// 	unit8(m.rng.Intn(129)),
	// 	0xFF,
	// }

	return nil
}

func (m *Image) encodePNG() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, m.Paletted); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}

// WriteTo writes captcha image in PNG format into the given writer
func (m *Image) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(m.encodePNG())
	return int64(n), err
}

func (m *Image) calculateSizes(width, height, ncount int) {
	// Goal: fit all digits inside the image
	var border int
	if width > height {
		border = height / 4
	} else {
		border = width / 4
	}
	_ = border
	// convert everything to floats for calculation
	// fw takes into account 1-dot spacing between digits
}
