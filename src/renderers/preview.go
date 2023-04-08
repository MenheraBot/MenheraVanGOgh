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

func RenderPreview(data *PreviewData) image.Image {
	var toReturn image.Image

	switch data.Type {
	case "table":
		image := utils.GetAsset("tables/" + data.Theme + ".png")
		ctx := gg.NewContextForImage(image)
		toReturn = ctx.Image()
	case "cards":
		ctx := gg.NewContext(360, 130)
		first := utils.GetAsset("cards/" + data.Theme + "/13.png")
		second := utils.GetAsset("cards/" + data.Theme + "/22.png")
		third := utils.GetAsset("cards/" + data.Theme + "/47.png")
		ctx.DrawImage(first, 0, 0)
		ctx.DrawImage(second, 120, 0)
		ctx.DrawImage(third, 240, 0)
		toReturn = ctx.Image()
	case "card_background":
		image := utils.GetAsset("card_backgrounds/" + data.Theme + ".png")
		ctx := gg.NewContextForImage(image)
		toReturn = ctx.Image()
	case "eb_background":
		image := utils.GetAsset("8ball/backgrounds/" + data.Theme + ".png")
		ctx := gg.NewContextForImage(image)
		toReturn = ctx.Image()
	case "eb_text_box":
		image := utils.GetAsset("8ball/text_boxes/" + data.Theme + ".png")
		ctx := gg.NewContextForImage(image)
		toReturn = ctx.Image()
	case "eb_menhera":
		ctx := gg.NewContext(1170, 440)
		first := utils.GetAsset("8ball/menheras/" + data.Theme + "/negative_1.png")
		second := utils.GetAsset("8ball/menheras/" + data.Theme + "/neutral_1.png")
		third := utils.GetAsset("8ball/menheras/" + data.Theme + "/positive_1.png")
		ctx.DrawImage(first, 0, 0)
		ctx.DrawImage(second, 391, 0)
		ctx.DrawImage(third, 782, 0)
		toReturn = ctx.Image()
	}

	return toReturn
}
