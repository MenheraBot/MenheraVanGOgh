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

func RenderTrisal(data *TrisalData) image.Image {
	ctx := gg.NewContext(768, 256)

	firstImage := utils.GetImageFromURL(data.UserOne, 256)
	secondImage := utils.GetImageFromURL(data.UserTwo, 256)
	thirdImage := utils.GetImageFromURL(data.UserThree, 256)

	ctx.DrawImage(firstImage, 0, 0)
	ctx.DrawImage(secondImage, 256, 0)
	ctx.DrawImage(thirdImage, 512, 0)

	return ctx.Image()
}
