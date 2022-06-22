package renderers

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type ShipData struct {
	LinkOne   string `json:"linkOne"`
	LinkTwo   string `json:"linkTwo"`
	ShipValue uint16 `json:"shipValue"`
}

func rainbowColorPercentage(percentage uint16) image.Image {
	ctx := gg.NewContext(456, 75)

	gradient := gg.NewLinearGradient(0, 0, 456, 75)
	gradient.AddColorStop(0.00, color.RGBA{255, 0, 0, 255})
	gradient.AddColorStop(1.0/6, color.RGBA{255, 100, 0, 255})
	gradient.AddColorStop(2.0/6, color.RGBA{255, 255, 0, 255})
	gradient.AddColorStop(3.0/6, color.RGBA{0, 128, 0, 255})
	gradient.AddColorStop(4.0/6, color.RGBA{0, 255, 255, 255})
	gradient.AddColorStop(5.0/6, color.RGBA{0, 0, 255, 255})
	gradient.AddColorStop(1.00, color.RGBA{84, 22, 180, 255})

	ctx.SetFillStyle(gradient)
	ctx.SetLineWidth(20)

	howMuchToFill := float64(percentage) / 100

	ctx.DrawRoundedRectangle(0, 0, 456*howMuchToFill, 75, 40)
	ctx.Fill()

	ctx.DrawRoundedRectangle(0, 0, 456, 75, 40)
	ctx.Stroke()

	return ctx.Image()
}

func RenderShip(data *ShipData, util utils.Utils) image.Image {
	ctx := gg.NewContext(512, 350)

	firstAvatar := util.ReadImageFromURL(data.LinkOne, 256, 256)
	secondAvatar := util.ReadImageFromURL(data.LinkTwo, 256, 256)
	shipLoadedImage := rainbowColorPercentage(data.ShipValue)

	ctx.DrawImage(firstAvatar, 0, 0)
	ctx.DrawImage(secondAvatar, 256, 0)
	ctx.DrawImage(shipLoadedImage, 20, 270)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 58)

	util.StrokeText(ctx, fmt.Sprint(data.ShipValue, "%"), 256, 330, 4, "#000", "#FFF")

	return ctx.Image()
}
