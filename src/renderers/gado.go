package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type GadoData struct {
	Image string `json:"image"` // 512
}

func RenderGado(data *GadoData) image.Image {
	ctx := gg.NewContext(1200, 526)

	userImage := utils.GetImageFromURL(data.Image, 455)
	gadoImage := utils.GetAsset("/images/gado.png")

	ctx.DrawImage(userImage, 695, 0)
	ctx.DrawImage(gadoImage, 0, 0)

	return ctx.Image()
}
