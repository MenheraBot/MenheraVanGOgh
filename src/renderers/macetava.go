package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type MacetavaData struct {
	Image               string `json:"image"`
	AuthorName          string `json:"authorName"`
	AuthorDiscriminator string `json:"authorDiscriminator"`
	AuthorImage         string `json:"authorImage"`
}

func rgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

func RenderMacetava(data *MacetavaData, util utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 882)

	userImage := util.GetImageFromURL(data.Image, 502, 573)
	userImageGayscale := rgbaToGray(userImage)
	userAvatar := util.GetImageFromURL(data.AuthorImage, 145, 145)
	macetavaImage := util.GetAsset("/images/macetava.png")

	ctx.DrawImage(userAvatar, 30, 18)
	ctx.DrawImage(userImage, 33, 305)
	ctx.DrawImage(userImageGayscale, 542, 305)

	ctx.DrawImage(macetavaImage, 0, 0)

	ctx.SetHexColor("#FFF")
	ctx.LoadFontFace(util.GetFontPath("Arial"), 48)
	util.StrokeText(ctx, data.AuthorName, 210, 85, 2, "#000", "#FFF", 0)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 38)
	ctx.SetHexColor("#86878C")
	ctx.DrawString(data.AuthorName+"#"+data.AuthorDiscriminator, 250, 145)

	return ctx.Image()
}
