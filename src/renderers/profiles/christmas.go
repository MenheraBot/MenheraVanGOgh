package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderChristmas(User *utils.UserData, I18n *utils.I18n, util *utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(67, 30, 950, 621)
	ctx.Fill()

	darkerColor := util.ShadeColor(baseColor, -15)
	ctx.SetHexColor(darkerColor)
	ctx.DrawRectangle(48, 465, 974, 187)
	ctx.Fill()

	ctx.DrawRoundedRectangle(370, 208, 557, 53, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	userAvatar := util.GetImageFromURL(User.Avatar, 250)

	ctx.DrawCircle(193, 133, 125)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 193, 133, 0.5, 0.5)
	ctx.ResetClip()

	backgroundImage := util.GetAsset("/profiles/christmas.png")

	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#FF0000")

	fontSize := 40

	if len(User.Username) > -18 {
		fontSize = 28
	}

	ctx.SetFontFace(*util.GetFont("Candy", float64(fontSize)))

	util.StrokeText(ctx, User.Username, 660, 100, 2, 0.5, 0, "#FFF")

	ctx.SetHexColor("#FF0000")
	ctx.DrawStringAnchored(User.Username, 660, 100, 0.5, 0)

	ctx.SetHexColor(util.GetCompatibleFontColor(baseColor))
	ctx.SetFontFace(*util.GetFont("Impact", 32))
	ctx.DrawStringWrapped(User.Info, 90, 540, 0, 1, 920, 1, 0)

	if User.Married {
		ctx.SetFontFace(*util.GetFont("Sans", 24))
		ctx.DrawStringWrapped(User.Marry.Username+" "+strings.Split(User.MarryDate, " ")[0], 400, 140, 0, 1, 600, 1, 0)
	}

	ctx.SetFontFace(*util.GetFont("Sans", 32))
	ctx.DrawStringAnchored(I18n.Mamado+": "+strconv.Itoa(User.Mamadas)+" "+I18n.Mamou+": "+strconv.Itoa(User.Mamou), 380, 243, 0, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 40))

	ctx.DrawStringWrapped(I18n.Usages+"   | "+strconv.Itoa(User.Votes)+" Upvotes", 90, 270, 0, 0, 920, 1, 0)

	util.DrawBadges(ctx, User, 80, 667)

	return ctx.Image()
}
