package utils

import (
	"image"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/patrickmn/go-cache"
	"golang.org/x/image/webp"
)

type Utils struct {
	default_image    image.Image
	images_cache     map[string]image.Image
	ttl_images_cache cache.Cache
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

func (util *Utils) GetFontPath(name string) string {
	workdir, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "../assets/fonts/Arial.ttf"
	}

	return workdir + "/assets/fonts/" + name + ".ttf"

}

func (util *Utils) GetResizedAsset(path string, w, h int) image.Image {
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

	resized := imaging.Resize(img, w, h, imaging.Lanczos)

	util.images_cache[path] = resized
	return resized
}

func (util *Utils) StrokeText(ctx *gg.Context, s string, x, y, n int, stroke, color string) {
	ctx.SetHexColor(stroke)
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := x + dx
			y := y + dy
			ctx.DrawStringAnchored(s, float64(x), float64(y), 0.5, 0)
		}
	}
	ctx.SetHexColor(color)
	ctx.DrawStringAnchored(s, float64(x), float64(y), 0.5, 0)
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

func (util *Utils) ReadImageFromURL(url string, x, y int) image.Image {
	var imagem image.Image = nil
	getImage, ok := util.ttl_images_cache.Get(url)

	if ok {
		return getImage.(image.Image)
	}

	res, err := http.Get(url)

	if err != nil {
		return util.default_image
	}

	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)

	if err != nil {
		webpData, errr := webp.Decode(res.Body)
		if errr != nil {
			return util.default_image
		}
		imagem = webpData

	} else {
		imagem = img
	}

	imagem = imaging.Fill(imagem, x, y, imaging.Center, imaging.NearestNeighbor)

	if !ok {
		util.ttl_images_cache.Add(url, imagem, cache.DefaultExpiration)
	}

	return imagem
}

func New() Utils {
	def := gg.NewContext(1, 1)
	def.SetRGB(255, 255, 0)
	def.Clear()

	return Utils{
		default_image:    def.Image(),
		images_cache:     make(map[string]image.Image),
		ttl_images_cache: *cache.New(time.Hour, 65*time.Minute),
	}
}
