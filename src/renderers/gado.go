package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type GadoData struct {
	Image string `json:"image"` // 512
}

func RenderGado(data *GadoData, util utils.Utils) image.Image {
	ctx := gg.NewContext(1200, 526)

	userImage := util.GetImageFromURL(data.Image, 455, 500)
	gadoImage := util.GetAsset("/images/gado.png")

	ctx.DrawImage(userImage, 695, 0)
	ctx.DrawImage(gadoImage, 0, 0)

	return ctx.Image()
}
