package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type TrisalData struct {
	UserOne   string `json:"userOne" form:"text"`
	UserTwo   string `json:"userTwo" form:"text"`
	UserThree string `json:"userThree" form:"text"`
}

func RenderTrisal(data *TrisalData, util utils.Utils) image.Image {
	ctx := gg.NewContext(768, 256)

	firstImage := util.ReadImageFromURL(data.UserOne, 256, 256)
	secondImage := util.ReadImageFromURL(data.UserTwo, 256, 256)
	thirdImage := util.ReadImageFromURL(data.UserThree, 256, 256)

	ctx.DrawImage(firstImage, 0, 0)
	ctx.DrawImage(secondImage, 256, 0)
	ctx.DrawImage(thirdImage, 512, 0)

	return ctx.Image()
}
