package blocky

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"strings"

	"github.com/notwithering/sgr"
)

// Encoder to encode image into printable terminal art.
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new encoder with the specified writer.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

// Encode encodes img into half-block art and writes it to the encoder.
func (e *Encoder) Encode(img image.Image) {
	halfBlock(e.w, img)
}

func halfBlock(w io.Writer, img image.Image) {
	bounds := img.Bounds()
	if bounds.Dx() <= 0 || bounds.Dy() <= 0 {
		return
	}

	var sb strings.Builder
	sb.Grow(bounds.Dx() * bounds.Dy() * 5) // preallocate estimated size to reduce reallocations

	// each line contains 2 pixels arranged vertically ▀▄
	for y := 0; y < bounds.Dy(); y += 2 {
		for x := 0; x < bounds.Dx(); x++ {
			// get upper and lower pixels
			upper := img.At(x, y)
			lower := color.RGBAModel.Convert(color.Transparent)
			if y+1 < bounds.Dy() {
				lower = img.At(x, y+1)
			}

			// convert both pixels to rgba
			upperRgba := color.RGBAModel.Convert(upper).(color.RGBA)
			lowerRgba := color.RGBAModel.Convert(lower).(color.RGBA)

			// determine the character and sgr codes for colors
			var char, code string
			if upperRgba.A > 0 && lowerRgba.A > 0 {
				char = "▄"
				code = sgr.BgColorRGB + fmt.Sprintf("%d;%d;%dm", upperRgba.R, upperRgba.G, upperRgba.B)
				code += sgr.FgColorRGB + fmt.Sprintf("%d;%d;%dm", lowerRgba.R, lowerRgba.G, lowerRgba.B)
			} else if upperRgba.A > 0 {
				char = "▀"
				code = sgr.FgColorRGB + fmt.Sprintf("%d;%d;%dm", upperRgba.R, upperRgba.G, upperRgba.B) + sgr.BgDefault
			} else if lowerRgba.A > 0 {
				char = "▄"
				code = sgr.BgDefault + sgr.FgColorRGB + fmt.Sprintf("%d;%d;%dm", lowerRgba.R, lowerRgba.G, lowerRgba.B)
			} else {
				char = " "
				code = sgr.BgDefault
			}

			// write the pair
			sb.WriteString(code)
			sb.WriteString(char)
		}

		// add newline only if its not the last line
		if y+2 < bounds.Dy() {
			sb.WriteString(sgr.Reset + "\r\n")
		}
	}

	sb.WriteString(sgr.Reset) // reset colors at the end
	fmt.Fprint(w, sb.String())
}

// EncodeToString encodes the specified image into half-block art and returns it as a string.
func EncodeToString(img image.Image) string {
	var sb strings.Builder
	halfBlock(&sb, img)
	return sb.String()
}
