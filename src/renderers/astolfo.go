package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type AstolfoData struct {
	Text string `json:"text"`
}

func RenderAstolfo(data *AstolfoData, util utils.Utils) image.Image {
	ctx := gg.NewContext(253, 330)

	astolfoImage := util.GetAsset("images/astolfo.png")

	ctx.DrawImage(astolfoImage, 0, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 20))

	ctx.SetHexColor("#000")

	ctx.DrawStringWrapped(data.Text, 70, 185, 0, 0, 160, 1, 1)

	return ctx.Image()
}
