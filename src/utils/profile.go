package utils

import (
	"image"
	"image/color"
	"math"
	"strconv"

	"github.com/fogleman/gg"
)

type BadgeData struct {
	Id       int    `json:"id"`
	ObtainAt string `json:"obtainAt"`
}

type MarryData struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
}

type UserData struct {
	Color         string      `json:"color"`
	Avatar        string      `json:"avatar"`
	Votes         int         `json:"votes"`
	Info          string      `json:"info"`
	Tag           string      `json:"tag"`
	Badges        []BadgeData `json:"badges"`
	HiddingBadges []int       `json:"hiddingBadges"`
	Username      string      `json:"username"`
	Mamadas       int         `json:"mamadas"`
	Mamou         int         `json:"mamou"`
	MarryDate     string      `json:"marryDate"`
	Marry         MarryData   `json:"marry"`
	Married       bool        `json:"married"`
}

type I18n struct {
	Aboutme string `json:"aboutme"`
	Mamado  string `json:"mamado"`
	Mamou   string `json:"mamou"`
	Usages  string `json:"usages"`
}

type ProfileData struct {
	User UserData `json:"user"`
	I18n I18n     `json:"i18n"`
	Type string   `json:"type"`
}

const badgeSize = 64

func includes(arr []int, target int) bool {
	var result bool = false

	for _, x := range arr {
		if x == target {
			result = true
			break
		}
	}

	return result
}

func (util *Utils) getUserBadges(user *UserData) []image.Image {
	var images []image.Image

	for _, badge := range user.Badges {
		if !includes(user.HiddingBadges, badge.Id) {
			img, ok := util.GetResizedAsset("badges/"+strconv.Itoa(badge.Id)+".png", badgeSize, badgeSize)
			if ok {
				images = append(images, img)
			}
		}
	}

	return images
}

func (util *Utils) DrawBadges(ctx *gg.Context, user *UserData, w, h int) {

	for i, badge := range util.getUserBadges(user) {
		ctx.DrawImage(badge, i*badgeSize+w, h)
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

func (util *Utils) GetColorLuminance(color color.RGBA) float64 {
	return float64(float64(0.299)*float64(color.R) + float64(0.587)*float64(color.G) + float64(0.114)*float64(color.B))
}

func (util *Utils) GetCompatibleFontColor(hex_color string) string {
	c, ok := parseHexColorFast(hex_color)
	if !ok {
		c = color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	}

	if math.Abs(util.GetColorLuminance(c)-util.GetColorLuminance(color.RGBA{R: 0, G: 0, B: 0, A: 255})) >= 128.0 {
		return "000000"
	} else {
		return "ffffff"
	}
}
