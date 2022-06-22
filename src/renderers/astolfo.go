package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type AstolfoData struct {
	Text string `json:"text" form:"text"`
}

func RenderAstolfo(data *AstolfoData, util utils.Utils) image.Image {
	ctx := gg.NewContext(253, 330)

	astolfoImage := util.GetAsset("images/astolfo.png")

	ctx.DrawImage(astolfoImage, 0, 0)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 20)

	ctx.SetHexColor("#000")

	util.FillText(ctx, data.Text, 72, 208, 160, 250, 20)

	return ctx.Image()
}
