/*
See page 22.
Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
number of cycles to 20 instead of the default 5. Use the strconv.Atoi func tion to convert the
string parameter into an integer. You can see its documentation with go doc strconv.Atoi.

Allows spcecification of Lissajous parameters. The following parameters and defaults shown below.
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
*/

package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("lisening on http://localhost:8000 ")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {

	parms := map[string]int{
		"cycles":  5,   // number of complete x oscillator revolutions
		"size":    100, // image canvas covers [-size..+size]
		"nframes": 64,  // number of animation frames
		"delay":   8,   // delay between frames in 10ms units
	}
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// fmt.Fprintf(w, "Host = %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		// If key is already in dictionary
		// convert value to int
		// and replace value in dictionary
		if _, ok := parms[k]; ok {
			vInt, intErr := strconv.Atoi(v[0])
			if intErr == nil {
				parms[k] = vInt
			}

		}
	}

	lissajous(w, parms)
}

func lissajous(out io.Writer, parms map[string]int) {
	const (
		// cycles  = 5     // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		// size    = 100   // image canvas covers [-size..+size]
		// nframes = 64    // number of animation frames
		// delay   = 8     // delay between frames in 10ms units
	)
	// set parameters from passed map
	cycles := parms["cycles"]
	size := parms["size"]
	nframes := parms["nframes"]
	delay := parms["delay"]

	// Set initial color
	myPalette := palette.WebSafe
	var colorIndex uint8 = uint8(len(myPalette) / 2)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, myPalette)
		for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		// cycle through palette drawing colors
		colorIndex++
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
