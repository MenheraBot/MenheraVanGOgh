package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type MacetavaData struct {
	Image               string `json:"image"` // 512
	AuthorName          string `json:"authorName"`
	AuthorDisplayName	string `json:authorDisplayName`
	AuthorImage         string `json:"authorImage"` // 128
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

func RenderMacetava(data *MacetavaData, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 882)

	userImage := utils.GetImageFromURL(data.Image, 573, 573, db)
	userImageGayscale := rgbaToGray(userImage)
	userAvatar := utils.GetImageFromURL(data.AuthorImage, 145, 145, db)
	macetavaImage := utils.GetAsset("/images/macetava.png")

	ctx.DrawImage(userAvatar, 30, 18)
	ctx.DrawImage(userImage, 33, 305)
	ctx.DrawImage(userImageGayscale, 542, 305)

	ctx.DrawImage(macetavaImage, 0, 0)

	ctx.SetHexColor("#FFF")
	ctx.SetFontFace(*utils.GetFont("Arial", 48))
	ctx.DrawStringAnchored(data.AuthorDisplayName, 210, 85, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Arial", 38))
	ctx.SetHexColor("#86878C")
	ctx.DrawString(data.AuthorName, 250, 145)

	return ctx.Image()
}
