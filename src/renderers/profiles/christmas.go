package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderChristmas(User *utils.UserData, I18n *utils.I18n, customEdits []string, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	if utils.GetProfileCustomization("useImage", customEdits) {
		backgroundImage := utils.GetImageFromURL(User.Image, 1080, 720, db)

		ctx.DrawImage(backgroundImage, 0, 0)
	} else {
		ctx.SetHexColor(User.Color)
		ctx.DrawRectangle(67, 30, 950, 621)
		ctx.Fill()

		darkerColor := utils.ShadeColor(User.Color, -15)
		ctx.SetHexColor(darkerColor)
		ctx.DrawRectangle(48, 465, 974, 187)
		ctx.Fill()

		ctx.DrawRoundedRectangle(370, 208, 557, 53, 20)
		ctx.FillPreserve()
		ctx.SetHexColor("#000")
		ctx.Stroke()
	}

	backgroundImage := utils.GetAsset("/profiles/christmas.png")

	ctx.DrawImage(backgroundImage, 0, 0)

	userAvatar := utils.GetImageFromURL(User.Avatar, 255, 255, db)
	ctx.DrawCircle(195, 133, 127.5)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 195, 133, 0.5, 0.5)
	ctx.ResetClip()

	ctx.SetFontFace(*utils.GetFont("Candy", float64(28)))
	utils.StrokeText(ctx, User.Username, 660, 130, 2, 0.5, 0, "#FFF")
	ctx.SetHexColor("#FF0000")
	ctx.DrawStringAnchored(User.Username, 660, 130, 0.5, 0)

	ctx.SetHexColor("#fff")
	ctx.SetFontFace(*utils.GetFont("Sans", 24))
	ctx.DrawStringAnchored(User.Title, 660, 80, 0.5, 0)

	if User.Married {
		ctx.DrawStringWrapped(User.MarryUsername+" "+strings.Split(User.MarryDate, " ")[0], 400, 160, 0, 1, 600, 1, 0)
	}

	ctx.SetHexColor(utils.GetCompatibleFontColor(User.Color))
	ctx.SetFontFace(*utils.GetFont("Impact", 32))
	ctx.DrawStringWrapped(User.Info, 90, 540, 0, 1, 920, 1, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 32))
	ctx.DrawStringAnchored(I18n.Mamado+": "+strconv.Itoa(int(User.Mamadas))+" "+I18n.Mamou+": "+strconv.Itoa(int(User.Mamou)), 380, 243, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 40))

	ctx.DrawStringWrapped(I18n.Usages+"   | "+strconv.Itoa(int(User.Votes))+" Upvotes", 90, 270, 0, 0, 920, 1, 0)

	utils.DrawBadges(ctx, User, 80, 656)

	return ctx.Image()
}
