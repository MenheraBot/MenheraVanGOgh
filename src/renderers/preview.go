package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type PreviewData struct {
	Type  string `json:"type"`
	Theme string `json:"theme"`
}

func RenderPreview(data *PreviewData, util *utils.Utils) image.Image {
	var toReturn image.Image

	switch data.Type {
	case "table":
		image, _ := util.GetResizedAsset("tables/"+data.Theme+".png", 630, 460)
		ctx := gg.NewContextForImage(image)
		toReturn = ctx.Image()
	case "cards":
		ctx := gg.NewContext(408, 187)
		first, _ := util.GetResizedAsset("cards/"+data.Theme+"/13.png", 136, 187)
		second, _ := util.GetResizedAsset("cards/"+data.Theme+"/22.png", 136, 187)
		third, _ := util.GetResizedAsset("cards/"+data.Theme+"/47.png", 136, 187)
		ctx.DrawImage(first, 0, 0)
		ctx.DrawImage(second, 0, 0)
		ctx.DrawImage(third, 0, 0)
		toReturn = ctx.Image()
	case "card_background":
		image, _ := util.GetResizedAsset("card_backgrounds/"+data.Theme+".png", 156, 242)
		ctx := gg.NewContextForImage(image)
		toReturn = ctx.Image()
	}

	return toReturn
}
