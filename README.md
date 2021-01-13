# Physarum Util

## Usage

```bash
go run main.go -size 1024 -config random  -path 1024_random/1024_random_%d.png -config "random"
```

## Montages

```bash

montage -geometry +2+2 -background black out1610375635894804.png out1610375711527622.png test.png

mkdir collages_2x2
for i in $(seq 1 5); do montage -geometry +0+0 -background black $(ls | shuf -n 4) collages_2x2/collage_$i.png; done

mkdir collages_3x3
for i in $(seq 1 5); do montage -geometry +0+0 -background black $(ls | shuf -n 9) collages_3x3/collage_$i.png; done

```

## Custom Configurations

```bash
# giraffe:
mkdir 1024_giraffe
go run main.go -size 1024 -path 1024_giraffe/1024_giraffe_%d.png -config "Config{1.2904124, 29.014423, 0.48230073, 0.5069842, 5, 0.1}, Config{0.00034416709, 0.091324545, 1.941707, 0.2719941, 5, 0.1}" -numOfExamples 20

# worm_layer:
mkdir 1024_worm_layer
go run main.go -size 1024 -path 1024_worm_layer/1024_worm_layer_%d.png -config "Config{0.46754807, 2.3493083, 0.73405844, 1.9963937, 5, 0.1},Config{0.8376874, 57.147926, 1.0695691, 0.8991064, 5, 0.1},Config{1.2544509, 46.368645, 1.9702296, 0.82022434, 5, 0.1}" -numOfExamples 20

# wheat_field:
mkdir 1024_wheat_field
go run main.go -size 1024 -path 1024_wheat_field/1024_wheat_field_%d.png -config "Config{0.058657303, 50.81802, 0.030490957, 0.38916337, 5, 0.1},Config{1.0893325, 29.964167, 1.4092915, 1.2136803, 5, 0.1},Config{1.6177789, 37.033314, 1.9790958, 0.9553208, 5, 0.1}" -numOfExamples 20

# river_tree:
mkdir 1024_river_tree
go run main.go -size 1024 -path 1024_river_tree/1024_river_tree_%d.png -config "Config{0.78554046, 48.176323, 0.39295194, 1.2483901, 5, 0.1},Config{0.7788923, 10.651063, 0.31162846, 0.899157, 5, 0.1},Config{0.1396707, 63.989666, 0.159785, 1.2148943, 5, 0.1},Config{0.1206861, 31.359587, 1.2821184, 0.94223183, 5, 0.1}" -numOfExamples 20


# lines_in_the_ocean:
mkdir 1024_lines_in_the_ocean
go run main.go -size 1024 -path 1024_lines_in_the_ocean/1024_lines_in_the_ocean_%d.png -config "Config{0.030216739, 42.50935, 0.5036654, 0.76111054, 5, 0.1},Config{1.1248671, 2.0523524, 0.82062626, 1.0933275, 5, 0.1}" -numOfExamples 20


# roots_in_the_sky:
mkdir 1024_roots_in_the_sky
go run main.go -size 1024 -path 1024_roots_in_the_sky/1024_roots_in_the_sky_%d.png -config "Config{1.3490353, 50.985237, 1.5146195, 1.5414921, 5, 0.1},Config{1.7305, 4.496922, 0.96046007, 0.25557137, 5, 0.1}" -numOfExamples 20


# roots_on_a_sponge:
mkdir 1024_roots_on_a_sponge
go run main.go -size 1024 -path 1024_roots_on_a_sponge/1024_roots_on_a_sponge_%d.png -config "Config{0.6162916, 12.934261, 0.9282113, 1.7465683, 5, 0.1},Config{0.6763068, 50.701546, 1.910328, 1.7586418, 5, 0.1}" -numOfExamples 20

# space:
mkdir 1024_space
go run main.go -size 1024 -path 1024_space/1024_space_%d.png -config "Config{1.3361955, 7.083375, 1.961523, 0.8821242, 5, 0.1},Config{0.9454778, 13.70873, 1.4255247, 1.1666366, 5, 0.1}" -numOfExamples 20

# cyclone:
mkdir 1024_cyclone
go run main.go -size 1024 -path 1024_cyclone/1024_cyclone_%d.png -config "Config{0.87946403, 42.838207, 0.97047323, 2.8447638, 5, 0.29681}, Config{1.7357124, 17.430664, 0.30490428, 2.1706762, 5, 0.27878627}" -numOfExamples 20

# dunes:
mkdir 1024_dunes
go run main.go -size 1024 -path 1024_dunes/1024_dunes_%d.png -config "Config{0.99931663, 44.21652, 1.9704952, 1.4215798, 5, 0.1580779},Config{1.9694986, 1.294038, 0.5384646, 1.1613986, 5, 0.21102181}" -numOfExamples 20

# dot_grid:
mkdir 1024_dot_grid
go run main.go -size 1024 -path 1024_dot_grid/1024_dot_grid_%d.png -config "Config{1.3433642, 49.39263, 0.91616887, 0.69644034, 5, 0.17888786},Config{0.0856143, 1.6695175, 1.8827246, 2.3155663, 5, 0.14249614},Config{0.7959472, 33.977413, 0.5246451, 2.2891424, 5, 0.22549233}" -numOfExamples 20

# untitled:
mkdir 1024_untitled
go run main.go -size 1024 -path 1024_untitled/1024_untitled_%d.png -config "Config{1.7433162, 56.586433, 0.45428953, 0.78228176, 5, 0.19172272},Config{1.8340914, 1.6538872, 1.4098115, 1.6714363, 5, 0.17746642},Config{0.0049473564, 13.269191, 0.033447478, 1.0102618, 5, 0.2197167},Config{0.37645763, 31.045816, 0.81319964, 2.5322618, 5, 0.10834738},Config{0.7355474, 14.832715, 0.2641479, 0.8953786, 5, 0.14977153}" -numOfExamples 20

# cool:
mkdir 1024_cool
go run main.go -size 1024 -path 1024_cool/1024_cool_%d.png -config "Config{1.4107815, 61.27741, 0.49201587, 1.3007548, 5, 0.49895996},Config{1.1534524, 13.299458, 0.48315683, 1.8219115, 5, 0.41845483},Config{0.31089303, 60.62575, 1.0241486, 0.39942655, 5, 0.4576149},Config{0.40245488, 27.844227, 1.9592205, 0.5504824, 5, 0.19568197},Config{1.227412, 1.7987814, 0.39546785, 1.2640203, 5, 0.14201605}" -numOfExamples 20



```
