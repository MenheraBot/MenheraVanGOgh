package utils

import (
	"image"
	"image/color"
	"math"
	"strconv"

	"github.com/fogleman/gg"
)

type MarryData struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
}

type UserData struct {
	Color         string    `json:"color"`
	Id            string    `json:"id"`
	Avatar        string    `json:"avatar"`
	Votes         uint16    `json:"votes"`
	Info          string    `json:"info"`
	Tag           string    `json:"tag"`
	Badges        []uint8   `json:"badges"`
	HiddingBadges []uint8   `json:"hiddingBadges"`
	Username      string    `json:"username"`
	Mamadas       int32     `json:"mamadas"`
	Mamou         int32     `json:"mamou"`
	MarryDate     string    `json:"marryDate"`
	Marry         MarryData `json:"marry"`
	Married       bool      `json:"married"`
}

type I18n struct {
	Aboutme string `json:"aboutme"`
	Mamado  string `json:"mamado"`
	Mamou   string `json:"mamou"`
	Usages  string `json:"usages"`
}

type ProfileData struct {
	User                UserData `json:"user"`
	StringedProfileData string   `json:"stringedProfileData"`
	I18n                I18n     `json:"i18n"`
	Type                string   `json:"type"`
}

const badgeSize = 64

func includes(arr []uint8, target uint8) bool {
	var result bool = false

	for _, x := range arr {
		if x == target {
			result = true
			break
		}
	}

	return result
}

func getUserBadges(user *UserData) []image.Image {
	var images []image.Image

	for _, badge := range user.Badges {
		if !includes(user.HiddingBadges, badge) {
			img, ok := GetResizedAsset("badges/"+strconv.Itoa(int(badge))+".png", badgeSize, badgeSize)
			if ok {
				images = append(images, img)
			}
		}
	}

	return images
}

func DrawBadges(ctx *gg.Context, user *UserData, w, h int) {
	for i, badge := range getUserBadges(user) {
		ctx.DrawImage(badge, i*badgeSize+w, h)
	}
}

func DrawVerticalBadges(ctx *gg.Context, user *UserData, w, h int) {
	for i, badge := range getUserBadges(user) {
		ctx.DrawImage(badge, w, i*badgeSize+h)
	}
}

func parseHexColorFast(s string) (c color.RGBA, ok bool) {
	c.A = 0xff
	ok = true

	if s[0] != '#' {
		return c, false
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		ok = false
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		ok = false
	}
	return
}

func GetColorLuminance(color color.RGBA) float64 {
	return float64(float64(0.299)*float64(color.R) + float64(0.587)*float64(color.G) + float64(0.114)*float64(color.B))
}

// https://stackoverflow.com/questions/3942878/how-to-decide-font-color-in-white-or-black-depending-on-background-color
func GetCompatibleFontColor(hex_color string) string {
	c, ok := parseHexColorFast(hex_color)
	if !ok {
		c = color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	}

	red := float64(c.R) / 255.0
	green := float64(c.G) / 255.0
	blue := float64(c.B) / 255.0

	if red <= 0.04045 {
		red /= 12.92
	} else {
		red = math.Pow(((red + 0.055) / 1.055), 2.4)
	}

	if green <= 0.04045 {
		green /= 12.92
	} else {
		green = math.Pow(((green + 0.055) / 1.055), 2.4)
	}

	if blue <= 0.04045 {
		blue /= 12.92
	} else {
		blue = math.Pow(((blue + 0.055) / 1.055), 2.4)
	}

	L := 0.2126*red + 0.7152*green + 0.0722*blue

	if L > 0.179 {
		return "#000000"
	}

	return "#ffffff"

}
