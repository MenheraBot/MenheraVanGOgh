package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type TrisalData struct {
	UserOne   string `json:"userOne"` // 256
	UserTwo   string `json:"userTwo"` // 256
	UserThree string `json:"userThree"` // 256
}

func RenderTrisal(data *TrisalData, util utils.Utils) image.Image {
	ctx := gg.NewContext(768, 256)

	firstImage := util.GetImageFromURL(data.UserOne, 256, 256)
	secondImage := util.GetImageFromURL(data.UserTwo, 256, 256)
	thirdImage := util.GetImageFromURL(data.UserThree, 256, 256)

	ctx.DrawImage(firstImage, 0, 0)
	ctx.DrawImage(secondImage, 256, 0)
	ctx.DrawImage(thirdImage, 512, 0)

	return ctx.Image()
}
