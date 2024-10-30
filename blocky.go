package blocky

import (
	"bytes"
	"image"
	"image/color"
	"io"
	"strconv"
	"strings"
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
	dx := bounds.Dx()
	dy := bounds.Dy()

	var buf bytes.Buffer

	// \x1b[38;2;255;255;255m\x1b[48;2;255;255;255m▀ (39 bytes)
	// 39 bytes per pixel pair / 2 pixels = ~20 bytes per pixel
	buf.Grow(dx * dy * 20) // preallocate estimated size to reduce reallocations

	// each line contains 2 pixels arranged vertically ▀▄
	for y := 0; y < dy; y += 2 {
		for x := 0; x < dx; x++ {
			// get upper pixel
			upperRgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)

			// get lower pixel, default to transparent if y+1 is out of bounds
			var lowerRgba color.RGBA
			if y+1 < dy {
				lowerRgba = color.RGBAModel.Convert(img.At(x, y+1)).(color.RGBA)
			}

			// write ansi code and character
			// putting each pixel channel value as string in a variable is actually slower
			// than just using strconv.Itoa a bazillion times, go optimizaation is weird
			if upperRgba.A > 0 && lowerRgba.A > 0 {
				buf.WriteString("\x1b[38;2;")
				buf.WriteString(strconv.Itoa(int(upperRgba.R)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(upperRgba.G)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(upperRgba.B)))
				buf.WriteByte('m')
				buf.WriteString("\x1b[48;2;")
				buf.WriteString(strconv.Itoa(int(lowerRgba.R)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(lowerRgba.G)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(lowerRgba.B)))
				buf.WriteString("m▀")
			} else if upperRgba.A > 0 {
				buf.WriteString("\x1b[38;2;")
				buf.WriteString(strconv.Itoa(int(upperRgba.R)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(upperRgba.G)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(upperRgba.B)))
				buf.WriteString("m\x1b[49m▀")
			} else if lowerRgba.A > 0 {
				buf.WriteString("\x1b[49m\x1b[38;2;")
				buf.WriteString(strconv.Itoa(int(lowerRgba.R)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(lowerRgba.G)))
				buf.WriteByte(';')
				buf.WriteString(strconv.Itoa(int(lowerRgba.B)))
				buf.WriteString("m▄")
			} else {
				buf.WriteString("\x1b[49m ")
			}
		}

		// add newline only if its not the last line
		if y+2 < dy {
			buf.WriteString("\x1b[0m\r\n")
		}
	}

	buf.WriteString("\x1b[0m") // reset colors at the end
	if w != nil {
		w.Write(buf.Bytes())
	}
}

// EncodeToString encodes the specified image into half-block art and returns it as a string.
func EncodeToString(img image.Image) string {
	var sb strings.Builder
	halfBlock(&sb, img)
	return sb.String()
}
