// Exericse 1.6: Modify the Lissajous program to produce images in multiple colors by adding
// more values to palette and then displaying them by changing the third argument of Set-
// ColorIndex in some interesting way.

// Generate output .gif using > ./1.6 > test.gif
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

// Modified var below from Chp 1 - lissajous
// var palette = []color.Color{color.White, color.Black}
// Modified var from exercise 1.4
// var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}
var myPalette = []color.Color{color.White, color.Black}

const (
	blackIndex = 0   // first color in palette
	redIndex   = 115 // next color in palette
)

// Initialize colors in palette
func initPalette() {
	/*
		for j := uint8(16); j <= 255; j += 16 {
			for k := uint8(16); k <= 255; k += 16 {
				for l := uint8(16); l <= 255; l += 16 {
					myPalette = append(myPalette, color.RGBA{j, k, l, uint8(255)})
				}
			}
		}
	*/

	// try this palette
	myPalette = palette.WebSafe
}

func main() {
	initPalette()

	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	var colorIndex uint8 = redIndex
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, myPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
			//colorIndex++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		colorIndex++
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
