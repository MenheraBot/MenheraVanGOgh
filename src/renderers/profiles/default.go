package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderDefault(User *utils.UserData, I18n *utils.I18n, util *utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	darker := util.ShadeColor(baseColor, -15)

	ctx.SetHexColor(darker)
	ctx.DrawRoundedRectangle(0, 0, 1080, 240, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	darkestThanTheDarkerColor := util.ShadeColor(darker, -10)

	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(0, 164, 1080, 75, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(860, 250, 210, 200, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	userAvatar := util.GetImageFromURL(User.Avatar, 250)

	ctx.SetHexColor("#000")
	ctx.DrawCircle(120, 120, 130)
	ctx.Fill()

	ctx.DrawCircle(120, 120, 120)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 120, 120, 0.5, 0.5)
	ctx.ResetClip()

	ctx.SetHexColor(util.GetCompatibleFontColor(darker))

	ctx.SetFontFace(*util.GetFont("Sans", 50))
	ctx.DrawStringWrapped(User.Tag, 255, 80, 0, 0.5, 650, 1, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 45))
	ctx.DrawStringAnchored("Upvotes", 860, 60, 0, 0)
	ctx.DrawStringAnchored(strconv.Itoa(User.Votes), 960, 120, 0.5, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 55))
	ctx.DrawStringAnchored(I18n.Aboutme, 20, 310, 0, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 40))
	ctx.DrawStringWrapped(User.Info, 20, 340, 0, 0, 860, 1, 0)

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 480, 1080, 720)
	ctx.Fill()

	ctx.SetHexColor(util.GetCompatibleFontColor(baseColor))

	ctx.SetFontFace(*util.GetFont("Sans", 50))
	ctx.DrawStringWrapped(I18n.Usages, 10, 550, 0, 0, 1070, 1, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 40))
	if User.Married {
		ringEmoji, _ := util.GetResizedAsset("badges/17.png", 64, 64)
		ctx.DrawStringAnchored(User.Marry.Tag+" | "+User.MarryDate, 80, 535, 0, 0)
		ctx.DrawImage(ringEmoji, 10, 490)
	}

	ctx.DrawStringAnchored(I18n.Mamado, 960, 290, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamou, 960, 380, 0.5, 0)

	ctx.DrawStringAnchored(strconv.Itoa(User.Mamadas), 960, 335, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(User.Mamou), 960, 425, 0.5, 0)

	util.DrawBadges(ctx, User, 230, 170)

	return ctx.Image()
}
