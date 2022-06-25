package utils

import (
	"fmt"
	"image"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/patrickmn/go-cache"
	"golang.org/x/image/font"
	"golang.org/x/image/webp"
)

type Utils struct {
	default_image    image.Image
	images_cache     map[string]image.Image
	ttl_images_cache cache.Cache
	fonts_cache      map[string]font.Face
}

func getFontPath(name string) string {
	workdir, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "../assets/fonts/Sans.ttf"
	}

	return workdir + "/assets/fonts/" + name + ".ttf"
}

func (util *Utils) GetResizedAsset(path string, w, h int) (image.Image, bool) {
	v, ok := util.images_cache[path+"-"+strconv.Itoa(w)+"-"+strconv.Itoa(h)]

	if ok {
		return v, true
	}

	workdir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return util.default_image, false
	}

	img_reader, err := os.Open(workdir + "/assets/" + path)
	if err != nil {
		log.Println(err)
		return util.default_image, false
	}

	defer img_reader.Close()

	img, _, err := image.Decode(img_reader)
	if err != nil {
		log.Println(err)
		return util.default_image, false
	}

	resized := imaging.Resize(img, w, h, imaging.Lanczos)

	util.images_cache[path+"-"+strconv.Itoa(w)+"-"+strconv.Itoa(h)] = resized
	return resized, true
}

func (util *Utils) GetFont(font string, points float64) *font.Face {
	v, ok := util.fonts_cache[font+"-"+strconv.Itoa(int(points))]

	if ok {
		return &v
	}

	face, err := gg.LoadFontFace(getFontPath(font), points)

	if err != nil {
		log.Panic(err)
	}

	util.fonts_cache[font+"-"+strconv.Itoa(int(points))] = face
	return &face
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

func (util *Utils) GetImageFromURL(url string, w int) image.Image {
	var imagem image.Image = nil
	getImage, ok := util.ttl_images_cache.Get(url + "-" + strconv.Itoa(w))

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

	imagem = imaging.Fill(imagem, w, w, imaging.Center, imaging.NearestNeighbor)

	util.ttl_images_cache.Add(url+"-"+strconv.Itoa(w), imagem, cache.DefaultExpiration)

	return imagem
}

func (util *Utils) ShadeColor(color string, percent float64) string {
	num, err := strconv.ParseInt(trimLeftChar(color), 16, 32)
	if err != nil {
		return color
	}

	amt := math.Round(2.55 * percent)

	R := (num >> 16) + int64(amt)
	G := (num >> 8 & 0x00FF) + int64(amt)
	B := (num & 0x0000FF) + int64(amt)

	if R < 255 {
		if R < 1 {
			R = 0
		}
	} else {
		R = 255
	}

	if G < 255 {
		if G < 1 {
			G = 0
		}
	} else {
		G = 255
	}

	if B < 255 {
		if B < 1 {
			B = 0
		}
	} else {
		B = 255
	}

	R *= 0x10000
	G *= 0x100

	return "#" + trimLeftChar(fmt.Sprintf("%x", 0x1000000+R+G+B))
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func (util *Utils) StrokeText(ctx *gg.Context, text string, x, y, size int, ax, ay float64, color string) {
	ctx.SetHexColor(color)
	for dy := -size; dy <= size; dy++ {
		for dx := -size; dx <= size; dx++ {
			if dx*dx+dy*dy >= size*size {
				// give it rounded corners
				continue
			}
			x := float64(x) + float64(dx)
			y := float64(y) + float64(dy)
			ctx.DrawStringAnchored(text, x, y, ax, ay)
		}
	}
}

func New() Utils {
	def := gg.NewContext(1, 1)
	def.SetRGB(255, 255, 0)
	def.Clear()

	return Utils{
		default_image:    def.Image(),
		images_cache:     make(map[string]image.Image),
		ttl_images_cache: *cache.New(30*time.Minute, 40*time.Minute),
		fonts_cache:      make(map[string]font.Face),
	}
}
