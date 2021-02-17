package main

import (
	"bufio"
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

func getModelSettingsString(model *physarum.Model, palette physarum.Palette, savePath string, zoom float64, particlesPowerOfTwo int) string {
	var out string
	out = ""
	out += fmt.Sprintf("%v ", savePath)
	out += fmt.Sprintf("-size %v ", model.W)
	out += fmt.Sprintf("-zoom %v ", zoom)
	out += fmt.Sprintf("-blurPasses %v ", model.BlurPasses)
	out += fmt.Sprintf("-blurRadius %v ", model.BlurRadius)
	out += fmt.Sprintf("-particlesPowerOfTwo %v ", particlesPowerOfTwo)
	out += fmt.Sprintf("-iterations %v ", model.Iteration)
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
	for _, c := range palette[:len(model.Configs)] {
		out += fmt.Sprintf("#%02X%02X%02X ", c.R, c.G, c.B)
	}
	out = out[:len(out)-1]
	out += "\"\n"
	return out
}

func main() {
	pathRaw := flag.String("path", "out_%d.png", "%d will be the time")
	size := flag.Int("size", 1024, "")
	configType := flag.String("config", "random", "")
	colorType := flag.String("color", "random", "")
	zoom := flag.Float64("zoom", 1, "")
	blurPasses := flag.Int("blurPasses", 2, "")
	blurRadius := flag.Int("blurRadius", 1, "")
	numOfExamples := flag.Int("numOfExamples", -1, "")
	particlesPowerOfTwo := flag.Int("particlesPowerOfTwo", 22, "")
	iterations := flag.Int("iterations", 400, "")
	logPath := flag.String("logPath", "", "")

	configLogPath := flag.String("configLogPath", "", "")
	configsLike := flag.String("configsLike", "", "")

	flag.Parse()

	// read log and use config from log
	if *configLogPath != "" && *configsLike != "" {
		file, err := os.Open(*configLogPath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			logConfigs := strings.Split(scanner.Text(), " ")
			if strings.Contains(logConfigs[0], *configsLike) {
				visitedFlags := ""
				flag.Visit(func(f *flag.Flag) {
					visitedFlags += f.Name + " "
				})
				logConfigsOtherSplit := strings.Split(scanner.Text(), "\"")
				if !strings.Contains(visitedFlags, "size ") {
					*size, err = strconv.Atoi(logConfigs[2])
				}
				if !strings.Contains(visitedFlags, "zoom ") {
					*zoom, err = strconv.ParseFloat(logConfigs[4], 10)
				}
				if !strings.Contains(visitedFlags, "blurPasses ") {
					*blurPasses, err = strconv.Atoi(logConfigs[6])
				}
				if !strings.Contains(visitedFlags, "blurRadius ") {
					*blurRadius, err = strconv.Atoi(logConfigs[8])
				}
				if !strings.Contains(visitedFlags, "particlesPowerOfTwo ") {
					*particlesPowerOfTwo, err = strconv.Atoi(logConfigs[10])
				}
				if !strings.Contains(visitedFlags, "iterations ") {
					*iterations, err = strconv.Atoi(logConfigs[12])
				}
				if !strings.Contains(visitedFlags, "config ") {
					*configType = logConfigsOtherSplit[1]
				}
				if !strings.Contains(visitedFlags, "particlesPowerOfTwo ") {
					*colorType = logConfigsOtherSplit[3]
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if *numOfExamples < 0 {
		*numOfExamples = 10000000
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < *numOfExamples; i++ {
		n, configs := getConfig(*configType)
		palette := getColor(*colorType)
		table := physarum.RandomAttractionTable(n)
		particles := 1 << *particlesPowerOfTwo
		model := physarum.NewModel(
			*size, *size, particles, *blurRadius, *blurPasses, float32(*size)/1024*float32(*zoom),
			configs, table)

		now := (time.Now().UTC().UnixNano() / 1e8) % 1e9
		savePath := fmt.Sprintf(*pathRaw, now)

		os.MkdirAll(path.Dir(savePath), os.ModePerm)
		one(model, *iterations, savePath, palette)

		modelString := getModelSettingsString(model, palette, savePath, *zoom, *particlesPowerOfTwo)

		if *logPath != "" {
			f, err := os.OpenFile(*logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(modelString); err != nil {
				log.Println(err)
			}
		}
	}
}
