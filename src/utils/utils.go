package utils

import (
	"fmt"
	"image"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/webp"
)

func defaultImage() image.Image {
	def := gg.NewContext(1, 1)
	def.SetRGB(255, 255, 0)
	def.Clear()

	return def.Image()
}

func getFontPath(name string) string {
	workdir, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return "../assets/fonts/Sans.ttf"
	}

	return workdir + "/assets/fonts/" + name + ".ttf"
}

func GetResizedAsset(path string, w, h int) (image.Image, bool) {
	workdir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return defaultImage(), false
	}

	img_reader, err := os.Open(workdir + "/assets/" + path)
	if err != nil {
		log.Println(err)
		return defaultImage(), false
	}

	defer img_reader.Close()

	img, _, err := image.Decode(img_reader)
	if err != nil {
		log.Println(err)
		return defaultImage(), false
	}

	resized := imaging.Resize(img, w, h, imaging.Lanczos)

	return resized, true
}

func GetFont(font string, points float64) *font.Face {
	face, err := gg.LoadFontFace(getFontPath(font), points)

	if err != nil {
		log.Panic(err)
	}

	return &face
}

func GetAsset(path string) image.Image {
	workdir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return defaultImage()
	}

	img_reader, err := os.Open(workdir + "/assets/" + path)
	if err != nil {
		log.Println(err)
		return defaultImage()
	}

	defer img_reader.Close()

	img, _, err := image.Decode(img_reader)
	if err != nil {
		log.Println(err)
		return defaultImage()
	}

	return img
}

func GetImageFromURL(url string, w int, h int, db *database.Database) image.Image {
	var imagem image.Image = nil

	if db.Client != nil {
		res, err := db.GetAvatar(url, w)
		if err == nil {
			imagem = res
			return imagem
		}
	}

	res, err := http.Get(url)

	if err != nil {
		return defaultImage()
	}

	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)

	if err != nil {
		webpData, errr := webp.Decode(res.Body)
		if errr != nil {
			return defaultImage()
		}
		imagem = webpData

	} else {
		imagem = img
	}

	imagem = imaging.Fill(imagem, w, h, imaging.Center, imaging.NearestNeighbor)

	if db.Client != nil {
		db.SetAvatar(url, w, imagem)
	}

	return imagem
}

func ShadeColor(color string, percent float64) string {
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

func StrokeText(ctx *gg.Context, text string, x, y, size int, ax, ay float64, color string) {
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
