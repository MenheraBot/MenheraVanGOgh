package utils

import (
	"image"
	"os"
	"log"

	"github.com/fogleman/gg"
)

type Utils struct {
	default_image image.Image
	images_cache  map[string]image.Image
}

func canFitHeightWise(ctx *gg.Context, lines []string, maxHeight, spacing int) bool {
	sum := 0
	for _, text := range lines {
		_, h := ctx.MeasureString(text)
		sum += int(h) + spacing
	}
	return sum < maxHeight
}

func (util *Utils) FillText(ctx *gg.Context, s string, x, y, width, height, spacing int) {
	lines := ctx.WordWrap(s, float64(width))
	var tbd []string

	for len(lines) > 0 && canFitHeightWise(ctx, append(tbd, lines[0]), height, spacing) {
		tbd = append(tbd, lines[0])
		lines = lines[1:]
	}

	currentY := y
	for _, text := range tbd {
		ctx.DrawString(text, float64(x), float64(currentY))
		currentY += spacing
	}
}

func  (util *Utils) GetFontPath(name string) string {
	workdir, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "../assets/fonts/Arial.ttf"
	}

	return workdir + "/assets/fonts/" + name + ".ttf"

}

func (util *Utils) GetAsset(path string) image.Image {
	v, ok := util.images_cache[path]

	if ok {
		return v
	}

	workdir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return util.default_image
	}

	img_reader, err := os.Open(workdir + "/assets/" + path)
	if err != nil {
		log.Println(err)
		return util.default_image
	}

	defer img_reader.Close()

	img, _, err := image.Decode(img_reader)
	if err != nil {
		log.Println(err)
		return util.default_image
	}

	util.images_cache[path] = img
	return img
}

func New() Utils {
	def := gg.NewContext(1, 1)
	def.SetRGB(255, 255, 0)
	def.Clear()

	return Utils{
		default_image: def.Image(),
		images_cache:  make(map[string]image.Image),
	}
}
