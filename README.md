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
		fmt.Printf("[%s] Encoded %dx%d in %f seconds.\n", time.Now().Format(time.DateTime), size, size, duration)
	}
}
```

</details>

```
[2024-10-30 16:54:37] Encoded 0x0 in 0.000002 seconds.
[2024-10-30 16:54:37] Encoded 100x100 in 0.004252 seconds.
[2024-10-30 16:54:37] Encoded 200x200 in 0.013289 seconds.
[2024-10-30 16:54:37] Encoded 300x300 in 0.015425 seconds.
[2024-10-30 16:54:37] Encoded 400x400 in 0.022112 seconds.
[2024-10-30 16:54:37] Encoded 500x500 in 0.032814 seconds.
[2024-10-30 16:54:37] Encoded 600x600 in 0.042764 seconds.
[2024-10-30 16:54:37] Encoded 700x700 in 0.056447 seconds.
[2024-10-30 16:54:37] Encoded 800x800 in 0.077747 seconds.
[2024-10-30 16:54:38] Encoded 900x900 in 0.100376 seconds.
[2024-10-30 16:54:38] Encoded 1000x1000 in 0.122180 seconds.
[2024-10-30 16:54:38] Encoded 1100x1100 in 0.161453 seconds.
[2024-10-30 16:54:38] Encoded 1200x1200 in 0.177685 seconds.
[2024-10-30 16:54:38] Encoded 1300x1300 in 0.217131 seconds.
[2024-10-30 16:54:39] Encoded 1400x1400 in 0.250649 seconds.
[2024-10-30 16:54:39] Encoded 1500x1500 in 0.272834 seconds.
[2024-10-30 16:54:40] Encoded 1600x1600 in 0.395836 seconds.
[2024-10-30 16:54:40] Encoded 1700x1700 in 0.352541 seconds.
[2024-10-30 16:54:41] Encoded 1800x1800 in 0.441595 seconds.
[2024-10-30 16:54:41] Encoded 1900x1900 in 0.494322 seconds.
[2024-10-30 16:54:42] Encoded 2000x2000 in 0.589988 seconds.
[2024-10-30 16:54:43] Encoded 2100x2100 in 0.614382 seconds.
[2024-10-30 16:54:44] Encoded 2200x2200 in 0.682691 seconds.
[2024-10-30 16:54:45] Encoded 2300x2300 in 0.653158 seconds.
[2024-10-30 16:54:46] Encoded 2400x2400 in 0.696960 seconds.
[2024-10-30 16:54:47] Encoded 2500x2500 in 0.819391 seconds.
[2024-10-30 16:54:48] Encoded 2600x2600 in 0.841819 seconds.
[2024-10-30 16:54:49] Encoded 2700x2700 in 0.908965 seconds.
[2024-10-30 16:54:50] Encoded 2800x2800 in 1.027381 seconds.
[2024-10-30 16:54:52] Encoded 2900x2900 in 1.028960 seconds.
[2024-10-30 16:54:53] Encoded 3000x3000 in 1.106413 seconds.
[2024-10-30 16:54:55] Encoded 3100x3100 in 1.245095 seconds.
[2024-10-30 16:54:57] Encoded 3200x3200 in 1.258750 seconds.
[2024-10-30 16:54:59] Encoded 3300x3300 in 1.345266 seconds.
[2024-10-30 16:55:01] Encoded 3400x3400 in 1.482347 seconds.
[2024-10-30 16:55:03] Encoded 3500x3500 in 1.507042 seconds.
[2024-10-30 16:55:05] Encoded 3600x3600 in 1.583149 seconds.
[2024-10-30 16:55:07] Encoded 3700x3700 in 1.710033 seconds.
[2024-10-30 16:55:09] Encoded 3800x3800 in 1.758081 seconds.
[2024-10-30 16:55:12] Encoded 3900x3900 in 1.851163 seconds.
[2024-10-30 16:55:15] Encoded 4000x4000 in 2.021311 seconds.
[2024-10-30 16:55:18] Encoded 4100x4100 in 2.086264 seconds.
[2024-10-30 16:55:21] Encoded 4200x4200 in 2.138776 seconds.
[2024-10-30 16:55:24] Encoded 4300x4300 in 2.371976 seconds.
[2024-10-30 16:55:27] Encoded 4400x4400 in 2.409068 seconds.
[2024-10-30 16:55:31] Encoded 4500x4500 in 2.555626 seconds.
[2024-10-30 16:55:34] Encoded 4600x4600 in 2.767860 seconds.
[2024-10-30 16:55:38] Encoded 4700x4700 in 2.794768 seconds.
[2024-10-30 16:55:42] Encoded 4800x4800 in 2.900084 seconds.
[2024-10-30 16:55:46] Encoded 4900x4900 in 2.834751 seconds.
[2024-10-30 16:55:50] Encoded 5000x5000 in 2.981798 seconds.
```