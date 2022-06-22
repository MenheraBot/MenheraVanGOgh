package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type PhiloData struct {
	Text string `json:"text"`
}

func RenderPhilo(data *PhiloData, util utils.Utils) image.Image {
	ctx := gg.NewContext(720, 720)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 58)

	ctx.SetHexColor("#FFF")

	util.FillText(ctx, data.Text, 0, 100, 720, 412, 50)

	philoImage := util.GetResizedAsset("images/philo.png", 720, 420)

	ctx.DrawImage(philoImage, 0, 300)

	return ctx.Image()
}
