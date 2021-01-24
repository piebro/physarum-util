package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"path"
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
	//blurPasses = 2
	zoomFactor = 1
)

func one(model *physarum.Model, iterations int, savePath string, palette physarum.Palette) {
	fmt.Println()
	fmt.Println(savePath)
	//fmt.Println(len(model.Particles), "particles")

	bar := pb.Full.Start(iterations)
	bar.SetWidth(80)
	for i := 0; i < iterations; i++ {
		bar.Increment()
		model.Step()
	}
	bar.Finish()

	im := physarum.Image(model.W, model.H, model.Data(), palette, 0, 0, 1/2.2)
	physarum.SavePNG(savePath, im, png.DefaultCompression)
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

func getColor(colorType string) physarum.Palette {
	if colorType == "random" {
		return physarum.RandomPalette()
	}

	colorType = strings.Replace(colorType, "#", "", -1)
	colorStrings := strings.Split(colorType, " ")
	pallete := make([]color.RGBA, 5)

	for i := 0; i < 5; i++ {
		colorInt, _ := strconv.ParseInt(colorStrings[i%len(colorStrings)], 16, 32)
		pallete[i] = physarum.HexColor(int(colorInt))
	}

	return pallete
}

func getModelSettingsString(model *physarum.Model, palette physarum.Palette, savePath string, zoom float64) string {
	var out string
	out = ""
	out += fmt.Sprintf("%v ", savePath)
	out += fmt.Sprintf("-size %v ", model.W)
	out += fmt.Sprintf("-zoom %v ", zoom)
	out += fmt.Sprintf("-blurPasses %v ", model.BlurPasses)
	out += "-config \""
	for _, c := range model.Configs {
		out += fmt.Sprintf("Config{%v, %v, %v, %v, %v, %v},",
			c.SensorAngle,
			c.SensorDistance,
			c.RotationAngle,
			c.StepDistance,
			c.DepositionAmount,
			c.DecayFactor)
	}
	out = out[:len(out)-1]
	out += "\" -color \""
	for _, c := range palette[:len(model.Configs)-1] {
		out += fmt.Sprintf("#%02X%02X%02X ", c.R, c.G, c.B)
	}
	out = out[:len(out)-1]
	out += "\"\n"
	return out
}

func main() {

	pathRaw := flag.String("path", "out%d.png", "%d will be the time")
	size := flag.Int("size", 1024, "")
	configType := flag.String("config", "random", "choices: cyclone, dunes, dot grid, untitled, cool")
	colorType := flag.String("color", "random", "")
	zoom := flag.Float64("zoom", 1, "")
	blurPasses := flag.Int("blurPasses", 2, "")
	numOfExamples := flag.Int("numOfExamples", -1, "")
	flag.Parse()

	if *numOfExamples < 0 {
		*numOfExamples = 10000000
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < *numOfExamples; i++ {
		n, configs := getConfig(*configType)
		palette := getColor(*colorType)
		table := physarum.RandomAttractionTable(n)
		model := physarum.NewModel(
			*size, *size, particles, blurRadius, *blurPasses, zoomFactor*float32(*size)/1024*float32(*zoom),
			configs, table)

		now := (time.Now().UTC().UnixNano() / 1e8) % 1e9
		savePath := fmt.Sprintf(*pathRaw, now)

		os.MkdirAll(path.Dir(savePath), os.ModePerm)
		one(model, iterations, savePath, palette)

		modelString := getModelSettingsString(model, palette, savePath, *zoom)
		f, err := os.OpenFile("generations.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(modelString); err != nil {
			log.Println(err)
		}
	}
}
