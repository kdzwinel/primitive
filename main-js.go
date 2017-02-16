package main

import (
	"log"
	"math/rand"
	"time"
	"encoding/base64"
	"image"
	"strings"

	_ "image/jpeg"

	"github.com/fogleman/primitive/primitive"
	"github.com/nfnt/resize"
	"github.com/gopherjs/gopherjs/js"
)

var (
	Input      string
	Config     shapeConfig
	Alpha      int
	InputSize  int
	OutputSize int
	Mode       int
	Workers    int
	Repeat     int
	Model      *primitive.Model
	Frame      int
)

type shapeConfig struct {
	Count  int
	Mode   int
	Alpha  int
	Repeat int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	js.Global.Set("primitive", map[string]interface{}{
		"init": initialize,
		"step": step,
	})
}

func initialize(o *js.Object) string {
	Input = o.Get("image").String()

	Alpha = o.Get("alpha").Int()
	InputSize = o.Get("inputSize").Int()
	OutputSize = o.Get("outputSize").Int()
	Mode = o.Get("mode").Int()
	Workers = 1
	Repeat = 0

	primitive.LogLevel = 2

	Config = shapeConfig{o.Get("primitivesCount").Int(), Mode, Alpha, Repeat}

	// seed random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// read input image
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(Input))
	input, _, err := image.Decode(reader)
	check(err)

	// scale down input image if needed
	size := uint(InputSize)
	if size > 0 {
		input = resize.Thumbnail(size, size, input, resize.Bilinear)
	}

	// determine background color
	var bg primitive.Color
	bg = primitive.MakeColor(primitive.AverageImageColor(input))

	// run algorithm
	Model = primitive.NewModel(input, bg, OutputSize, Workers)
	primitive.Log(1, "%d: t=%.3f, score=%.6f\n", 0, 0.0, Model.Score)

	Frame = 0

	// write output image
	return Model.SVG()
}

func step() string {
	Frame++
	t := time.Now()
	n := Model.Step(primitive.ShapeType(Config.Mode), Config.Alpha, Config.Repeat)
	nps := primitive.NumberString(float64(n) / time.Since(t).Seconds())
	primitive.Log(1, "%d: score=%.6f, n=%d, n/s=%s\n", Frame, Model.Score, n, nps)

	return Model.SVG();
}
