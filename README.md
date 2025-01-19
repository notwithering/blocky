# blocky

semi-fast half block art encoder for go

## examples

```go
// encode image to terminal
blocky.NewEncoder(os.Stdout).Encode(img)

// encode and put in variable
var art string = blocky.EncodeToString(img)
```

## benchmark

<details>
	<summary>code</summary>

```go
package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/notwithering/blocky"
)

func main() {
	encoder := blocky.NewEncoder(nil)

	for size := 0; size <= 5000; size += 100 {
		img := image.NewRGBA(image.Rect(0, 0, size, size))
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				var c color.Color

				switch (x + y) % 2 {
				case 0:
					c = color.White
				case 1:
					c = color.Black
				}

				img.Set(x, y, c)
			}
		}

		start := time.Now()

		encoder.Encode(img)

		duration := time.Since(start).Seconds()
		fmt.Printf("Encoded %dx%d in %f seconds (%.2f FPS).\n", size, size, duration, 1.0/duration)
	}
}
```

</details>

```
Encoded 0x0 in 0.000000 seconds (2217294.90 FPS).
Encoded 100x100 in 0.000528 seconds (1893.21 FPS).
Encoded 200x200 in 0.002077 seconds (481.49 FPS).
Encoded 300x300 in 0.005246 seconds (190.61 FPS).
Encoded 400x400 in 0.009811 seconds (101.92 FPS).
Encoded 500x500 in 0.012870 seconds (77.70 FPS).
Encoded 600x600 in 0.017968 seconds (55.65 FPS).
Encoded 700x700 in 0.025507 seconds (39.20 FPS).
Encoded 800x800 in 0.032529 seconds (30.74 FPS).
Encoded 900x900 in 0.040987 seconds (24.40 FPS).
Encoded 1000x1000 in 0.050666 seconds (19.74 FPS).
Encoded 1100x1100 in 0.060712 seconds (16.47 FPS).
Encoded 1200x1200 in 0.071966 seconds (13.90 FPS).
Encoded 1300x1300 in 0.084150 seconds (11.88 FPS).
Encoded 1400x1400 in 0.099162 seconds (10.08 FPS).
Encoded 1500x1500 in 0.111873 seconds (8.94 FPS).
Encoded 1600x1600 in 0.126627 seconds (7.90 FPS).
Encoded 1700x1700 in 0.143284 seconds (6.98 FPS).
Encoded 1800x1800 in 0.162943 seconds (6.14 FPS).
Encoded 1900x1900 in 0.181991 seconds (5.49 FPS).
Encoded 2000x2000 in 0.200929 seconds (4.98 FPS).
Encoded 2100x2100 in 0.217417 seconds (4.60 FPS).
Encoded 2200x2200 in 0.241479 seconds (4.14 FPS).
Encoded 2300x2300 in 0.263161 seconds (3.80 FPS).
Encoded 2400x2400 in 0.282993 seconds (3.53 FPS).
Encoded 2500x2500 in 0.312574 seconds (3.20 FPS).
Encoded 2600x2600 in 0.345916 seconds (2.89 FPS).
Encoded 2700x2700 in 0.365675 seconds (2.73 FPS).
Encoded 2800x2800 in 0.396629 seconds (2.52 FPS).
Encoded 2900x2900 in 0.429525 seconds (2.33 FPS).
Encoded 3000x3000 in 0.455866 seconds (2.19 FPS).
Encoded 3100x3100 in 0.483610 seconds (2.07 FPS).
Encoded 3200x3200 in 0.522303 seconds (1.91 FPS).
Encoded 3300x3300 in 0.544287 seconds (1.84 FPS).
Encoded 3400x3400 in 0.589076 seconds (1.70 FPS).
Encoded 3500x3500 in 0.625883 seconds (1.60 FPS).
Encoded 3600x3600 in 0.654770 seconds (1.53 FPS).
Encoded 3700x3700 in 0.698659 seconds (1.43 FPS).
Encoded 3800x3800 in 0.730801 seconds (1.37 FPS).
Encoded 3900x3900 in 0.767212 seconds (1.30 FPS).
Encoded 4000x4000 in 0.812765 seconds (1.23 FPS).
Encoded 4100x4100 in 0.852297 seconds (1.17 FPS).
Encoded 4200x4200 in 0.881649 seconds (1.13 FPS).
Encoded 4300x4300 in 0.921869 seconds (1.08 FPS).
Encoded 4400x4400 in 0.959685 seconds (1.04 FPS).
Encoded 4500x4500 in 1.033940 seconds (0.97 FPS).
Encoded 4600x4600 in 1.061284 seconds (0.94 FPS).
Encoded 4700x4700 in 1.127033 seconds (0.89 FPS).
Encoded 4800x4800 in 1.209544 seconds (0.83 FPS).
Encoded 4900x4900 in 1.223532 seconds (0.82 FPS).
Encoded 5000x5000 in 1.267292 seconds (0.79 FPS).
```
