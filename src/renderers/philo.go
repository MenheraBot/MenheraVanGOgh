package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/nfnt/resize"
)

type PhiloData struct {
	Text string `json:"text" form:"text"`
}

func RenderPhilo(data *PhiloData, util utils.Utils) image.Image {
	ctx := gg.NewContext(720, 720)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 58)

	ctx.SetHexColor("#FFF")

	util.FillText(ctx, data.Text, 0, 100, 720, 412, 50)

	philoImage := util.GetAsset("images/philo.png")

	ctx.DrawImage(resize.Resize(720, 420, philoImage, resize.Lanczos3), 0, 300)

	return ctx.Image()
}
