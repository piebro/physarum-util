package main

import (
	"flag"
	"fmt"
	"image/png"
	"math/rand"
	_ "net/http/pprof"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/fogleman/physarum/pkg/physarum"
)

const (
	width      = 1024
	height     = 1024
	particles  = 1 << 22
	iterations = 400
	blurRadius = 1
	blurPasses = 2
	zoomFactor = 1
)

func one(model *physarum.Model, iterations int, path string) {
	fmt.Println()
	fmt.Println(path)
	//fmt.Println(len(model.Particles), "particles")

	for _, c := range model.Configs {
		fmt.Printf("Config{%v, %v, %v, %v, %v, %v},",
			c.SensorAngle,
			c.SensorDistance,
			c.RotationAngle,
			c.StepDistance,
			c.DepositionAmount,
			c.DecayFactor)
	}
	fmt.Println()

	bar := pb.Full.Start(iterations)
	bar.SetWidth(80)
	for i := 0; i < iterations; i++ {
		bar.Increment()
		model.Step()
	}
	bar.Finish()

	palette := physarum.RandomPalette()
	im := physarum.Image(model.W, model.H, model.Data(), palette, 0, 0, 1/2.2)
	physarum.SavePNG(path, im, png.DefaultCompression)
}

func getConfig(configType string) (int, []physarum.Config) {
	if configType == "random" {
		n := 2 + rand.Intn(4)
		return n, physarum.RandomConfigs(n)

	}

	count := strings.Count(configType, "Config")
	r, _ := regexp.Compile(`(\d*\.\d+|\d+)`)
	c := r.FindAllString(configType, -1)
	configs := make([]physarum.Config, count)
	for i := 0; i < count; i++ {
		SensorAngle, _ := strconv.ParseFloat(c[i*6+0], 32)
		SensorDistance, _ := strconv.ParseFloat(c[i*6+1], 32)
		RotationAngle, _ := strconv.ParseFloat(c[i*6+2], 32)
		StepDistance, _ := strconv.ParseFloat(c[i*6+3], 32)
		DepositionAmount, _ := strconv.ParseFloat(c[i*6+4], 32)
		DecayFactor, _ := strconv.ParseFloat(c[i*6+5], 32)
		configs[i] = physarum.Config{
			SensorAngle:      float32(SensorAngle),
			SensorDistance:   float32(SensorDistance),
			RotationAngle:    float32(RotationAngle),
			StepDistance:     float32(StepDistance),
			DepositionAmount: float32(DepositionAmount),
			DecayFactor:      float32(DecayFactor),
		}
	}
	return count, configs

}

func main() {

	configType := flag.String("config", "random", "choices: cyclone, dunes, dot grid, untitled, cool")
	size := flag.Int("size", 1024, "")
	pathRaw := flag.String("path", "out%d.png", "%d will be the time")
	numOfExamples := flag.Int("numOfExamples", -1, "")
	flag.Parse()

	if *numOfExamples < 0 {
		*numOfExamples = 10000000
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < *numOfExamples; i++ {

		n, configs := getConfig(*configType)
		table := physarum.RandomAttractionTable(n)
		model := physarum.NewModel(
			*size, *size, particles, blurRadius, blurPasses, zoomFactor*float32(*size)/1024,
			configs, table)
		start := time.Now()

		now := (time.Now().UTC().UnixNano() / 1e8) % 1e6
		path := fmt.Sprintf(*pathRaw, now)
		one(model, iterations, path)
		fmt.Println(time.Since(start))
	}
}
