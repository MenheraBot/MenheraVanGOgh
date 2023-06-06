package renderers

import (
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type TrisalData struct {
	UserOne   string `json:"userOne"`   // 256
	UserTwo   string `json:"userTwo"`   // 256
	UserThree string `json:"userThree"` // 256
}

func RenderTrisal(data *TrisalData, db *database.Database) image.Image {
	ctx := gg.NewContext(768, 256)

	firstImage := utils.GetImageFromURL(data.UserOne, 256, 256, db)
	secondImage := utils.GetImageFromURL(data.UserTwo, 256, 256, db)
	thirdImage := utils.GetImageFromURL(data.UserThree, 256, 256, db)

	ctx.DrawImage(firstImage, 0, 0)
	ctx.DrawImage(secondImage, 256, 0)
	ctx.DrawImage(thirdImage, 512, 0)

	return ctx.Image()
}
