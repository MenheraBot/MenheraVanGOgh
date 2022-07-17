package renderers

import (
	"image"
	"image/color"
	"strconv"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type PokerUserData struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Theme  string `json:"theme"`
	Fold   bool   `json:"fold"`
	Chips  int    `json:"chips"`
	Dealer bool   `json:"dealer"`
}
type PokerTableData struct {
	ComunityCards []uint8         `json:"cards"`
	Users         []PokerUserData `json:"users"`
	Pot           int             `json:"pot"`
}

type PokerHandData struct {
	Cards []uint8 `json:"cards"`
	Theme string  `json:"theme"`
}

func RenderPokerHand(data *PokerHandData) image.Image {
	ctx := gg.NewContext(272, 187)

	firstImage, _ := utils.GetResizedAsset("cards/"+data.Theme+"/"+strconv.Itoa(int(data.Cards[0]))+".png", 136, 187)
	secondImage, _ := utils.GetResizedAsset("cards/"+data.Theme+"/"+strconv.Itoa(int(data.Cards[1]))+".png", 136, 187)

	ctx.DrawImage(firstImage, 0, 0)
	ctx.DrawImage(secondImage, 137, 0)
	return ctx.Image()
}

var avatarLocations = [9][2]uint16{{670, 70}, {860, 240}, {800, 460}, {580, 530}, {330, 530}, {130, 460}, {70, 240}, {240, 70}, {455, 70}}
var chipLocations = [9][2]uint16{{670, 140}, {720, 230}, {660, 390}, {510, 390}, {330, 390}, {170, 320}, {150, 220}, {240, 130}, {405, 130}}

func limitString(s string, size uint) string {
	if len(s) > int(size) {
		return s[:size]
	}
	return s
}

func drawAvatar(ctx *gg.Context, avatar image.Image, x, y uint16, m bool) {
	ctx.SetHexColor("#787878")
	ctx.DrawCircle(float64(x), float64(y), 65)
	ctx.Fill()

	if m {
		ctx.SetHexColor("#e64ce3")
	} else {
		ctx.SetHexColor("#b6b3b3")
	}

	ctx.DrawCircle(float64(x), float64(y), 63)
	ctx.Fill()

	ctx.DrawCircle(float64(x), float64(y), 60)
	ctx.Clip()
	ctx.DrawImageAnchored(avatar, int(x), int(y), 0.5, 0.5)
	ctx.ResetClip()
}

func getUserChipsImage(chips int) image.Image {

	var image image.Image

	if chips < 10000 {
		image, _ = utils.GetResizedAsset("poker/less_chips.png", 50, 50)
	} else if chips < 100000 {
		image, _ = utils.GetResizedAsset("poker/medium_chips.png", 60, 60)
	} else if chips < 500000 {
		image, _ = utils.GetResizedAsset("poker/lots_chips.png", 80, 80)
	} else {
		image, _ = utils.GetResizedAsset("poker/millions_chips.png", 80, 80)
	}

	return image
}

func RenderPokerTable(data *PokerTableData) image.Image {
	ctx := gg.NewContext(930, 600)

	background, _ := utils.GetResizedAsset("tables/green.png", 930, 600)
	ctx.DrawImage(background, 0, 0)

	tableImage, _ := utils.GetResizedAsset("tables/rounded.png", 830, 460)
	ctx.DrawImage(tableImage, 50, 70)

	startCardW := (450 - 31*len(data.ComunityCards))

	for i, card := range data.ComunityCards {
		cardImage, _ := utils.GetResizedAsset("cards/default/"+strconv.Itoa(int(card))+".png", 62, 76)
		ctx.DrawImage(cardImage, startCardW+(70*i), 250)
	}

	ctx.SetFontFace(*utils.GetFont("Arial", 16))

	for i, user := range data.Users {
		userAvatar := utils.GetImageFromURL(user.Avatar, 120)
		drawAvatar(ctx, userAvatar, avatarLocations[i][0], avatarLocations[i][1], false)

		if !user.Fold {
			userCardBackground, _ := utils.GetResizedAsset("card_backgrounds/"+user.Theme+".png", 37, 51)
			ctx.DrawImage(userCardBackground, int(avatarLocations[i][0])+10, int(avatarLocations[i][1])+10)
			ctx.DrawImage(userCardBackground, int(avatarLocations[i][0])+15, int(avatarLocations[i][1])+15)
		} else {
			ctx.SetColor(color.RGBA{R: 0, G: 0, B: 0, A: 200})
			ctx.DrawCircle(float64(avatarLocations[i][0]), float64(avatarLocations[i][1]), 63)
			ctx.Fill()
		}

		var anchorX float64 = 0
		var toLeft uint16 = 75

		if i == 1 && len(user.Name) > 15 {
			anchorX = 0.3
		}

		if i == 6 {
			toLeft = 70
			if len(user.Name) > 15 {
				anchorX = -0.3
			}
		}

		image := getUserChipsImage(user.Chips)
		ctx.DrawImage(image, int(chipLocations[i][0]), int(chipLocations[i][1]))

		nameSize, _ := ctx.MeasureString(limitString(user.Name, 20))
		numberSize, _ := ctx.MeasureString(strconv.Itoa(user.Chips))

		var textSize = nameSize

		if numberSize > nameSize {
			textSize = numberSize
		}

		ctx.SetColor(color.RGBA{R: 0, G: 0, B: 0, A: 180})
		ctx.DrawRoundedRectangle(float64(avatarLocations[i][0])-textSize/2-10, float64(avatarLocations[i][1]-70), textSize+10, 36, 10)
		ctx.Fill()

		if user.Fold {
			ctx.SetHexColor("#FF0000")
		} else {
			ctx.SetHexColor("#FFF")
		}

		ctx.DrawStringWrapped(ctx.WordWrap(limitString(user.Name, 20), 140)[0], float64(avatarLocations[i][0]-toLeft), float64(avatarLocations[i][1]-70), anchorX, 0, 140, 1, 1)
		ctx.SetHexColor("#FFFF00")
		ctx.DrawStringWrapped(strconv.Itoa(user.Chips), float64(avatarLocations[i][0]-toLeft), float64(avatarLocations[i][1]-56), anchorX, 0, 140, 1, 1)

		if user.Dealer {
			dealerBotton, _ := utils.GetResizedAsset("poker/dealer.png", 60, 60)
			ctx.DrawImageAnchored(dealerBotton, int(avatarLocations[i][0]-45), int(avatarLocations[i][1]+35), 0.5, 0.5)
		}
	}

	menheraAvatar, _ := utils.GetResizedAsset("poker/headphone.png", 120, 120)
	drawAvatar(ctx, menheraAvatar, avatarLocations[8][0], avatarLocations[8][1], true)

	potImage := getUserChipsImage(data.Pot)

	ctx.DrawImage(potImage, int(chipLocations[8][0]), int(chipLocations[8][1]))

	ctx.SetColor(color.RGBA{R: 0, G: 0, B: 0, A: 180})
	ctx.DrawRoundedRectangle(float64(startCardW), 200, float64(len(data.ComunityCards)*62)+31, 40, 5)
	ctx.Fill()

	ctx.SetHexColor("#FFFF00")
	ctx.SetFontFace(*utils.GetFont("Arial", 40))

	ctx.DrawStringAnchored("Pot: "+strconv.Itoa(data.Pot), float64(startCardW), 200, 0, 0.86)

	return ctx.Image()
}
