package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type PhiloData struct {
	Text string `json:"text"`
}

func RenderPhilo(data *PhiloData) image.Image {
	ctx := gg.NewContext(720, 720)

	ctx.SetFontFace(*utils.GetFont("Sans", 58))

	ctx.SetHexColor("#FFF")

	ctx.DrawStringWrapped(data.Text, 0, 0, 0, 0, 720, 1, 1)

	philoImage, _ := utils.GetResizedAsset("images/philo.png", 720, 420)

	ctx.DrawImage(philoImage, 0, 300)

	return ctx.Image()
}
