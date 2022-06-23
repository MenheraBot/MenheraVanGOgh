package renderers

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type ShipData struct {
	LinkOne   string `json:"linkOne"` // 256
	LinkTwo   string `json:"linkTwo"` // 256
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

	firstAvatar := util.GetImageFromURL(data.LinkOne, 256)
	secondAvatar := util.GetImageFromURL(data.LinkTwo, 256)
	shipLoadedImage := rainbowColorPercentage(data.ShipValue)

	ctx.DrawImage(firstAvatar, 0, 0)
	ctx.DrawImage(secondAvatar, 256, 0)
	ctx.DrawImage(shipLoadedImage, 20, 270)

	ctx.SetFontFace(*util.GetFont("Sans", 58))

	text := fmt.Sprint(data.ShipValue, "%")

	util.StrokeText(ctx, text, 256, 330, 2, 0.5, 0, "#000")

	ctx.SetRGBA255(255, 255, 255, 255)
	ctx.DrawStringAnchored(text, 256, 330, 0.5, 0)

	return ctx.Image()
}
