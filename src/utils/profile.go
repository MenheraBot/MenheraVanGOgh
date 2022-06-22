package utils

import (
	"image"
	"strconv"

	"github.com/fogleman/gg"
)

type BadgeData struct {
	Id       int `json:"id"`
	ObtainAt int `json:"obtainAt"`
}

type MarryData struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
}

type UserData struct {
	Cor           string      `json:"cor"`
	Avatar        string      `json:"avatar"`
	Votos         int         `json:"votos"`
	Nota          string      `json:"nota"`
	Tag           string      `json:"tag"`
	FlagsArray    []string    `json:"flagsArray"`
	Casado        string      `json:"casado"`
	VoteCooldown  int         `json:"voteCooldown"`
	Badges        []BadgeData `json:"badges"`
	HiddingBadges []int       `json:"hiddingBadges"`
	Username      string      `json:"username"`
	Data          string      `json:"data"`
	Mamadas       int         `json:"mamadas"`
	Mamou         int         `json:"mamou"`
	Marry         MarryData   `json:"marry"`
	Married       bool        `json:"married"`
}

type Cmds struct {
	Count int `json:"count"`
}

type Array struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type UsageCommands struct {
	Cmds  Cmds    `json:"cmds"`
	Array []Array `json:"array"`
}

type I18n struct {
	Aboutme string `json:"aboutme"`
	Mamado  string `json:"mamado"`
	Mamou   string `json:"mamou"`
	Zero    string `json:"zero"`
	Um      string `json:"um"`
	Dois    string `json:"dois"`
	Tres    string `json:"tres"`
}

type ProfileData struct {
	User          UserData      `json:"user"`
	UsageCommands UsageCommands `json:"usageCommands"`
	I18n          I18n          `json:"i18n"`
	Type          string        `json:"type"`
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
