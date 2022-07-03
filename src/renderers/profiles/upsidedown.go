package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderUpsideDown(User *utils.UserData, I18n *utils.I18n) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	ctx.SetHexColor(utils.GetCompatibleFontColor(baseColor))
	ctx.SetFontFace(*utils.GetFont("Sans", 45))
	ctx.DrawStringAnchored(I18n.Aboutme, 20, 250, 0, 0.5)

	ctx.SetFontFace(*utils.GetFont("Sans", 40))
	ctx.DrawStringWrapped(User.Info, 20, 340, 0, 0.5, 860, 1, 0)

	darker := utils.ShadeColor(baseColor, 25)

	ctx.SetHexColor(darker)
	ctx.DrawRoundedRectangle(0, 480, 1080, 240, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	darkestThanTheDarkerColor := utils.ShadeColor(darker, 10)

	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(0, 480, 1080, 75, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(870, 250, 200, 200, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	ctx.SetHexColor(utils.GetCompatibleFontColor(darkestThanTheDarkerColor))
	ctx.SetFontFace(*utils.GetFont("Sans", 50))
	ctx.DrawStringWrapped(User.Tag, 255, 635, 0, 0.5, 650, 1, 0)

	ctx.SetHexColor("#000")
	ctx.DrawCircle(960, 600, 130)
	ctx.Fill()

	userAvatar := utils.GetImageFromURL(User.Avatar, 250)
	ctx.DrawCircle(960, 600, 120)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 960, 600, 0.5, 0.5)
	ctx.ResetClip()
	ctx.SetHexColor(utils.ShadeColor(baseColor, 60))
	ctx.DrawRoundedRectangle(0, 555, 240, 165, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	ctx.SetHexColor(utils.GetCompatibleFontColor(darker))
	ctx.SetFontFace(*utils.GetFont("Sans", 45))
	ctx.DrawStringAnchored("Upvotes", 20, 620, 0, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Votes)), 120, 690, 0.5, 0)

	ctx.SetHexColor(utils.GetCompatibleFontColor(baseColor))
	ctx.SetFontFace(*utils.GetFont("Sans", 50))
	ctx.DrawStringWrapped(I18n.Usages, 10, 70, 0, 0.5, 1070, 1, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 40))

	if User.Married {
		ringEmoji, _ := utils.GetResizedAsset("badges/17.png", 64, 64)
		ctx.DrawStringAnchored(User.Marry.Tag+" | "+User.MarryDate, 80, 210, 0, 0)
		ctx.DrawImage(ringEmoji, 10, 165)
	}

	ctx.SetHexColor(utils.GetCompatibleFontColor(darkestThanTheDarkerColor))

	ctx.DrawStringAnchored(I18n.Mamado, 970, 290, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamou, 970, 380, 0.5, 0)

	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamadas)), 965, 335, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamou)), 965, 425, 0.5, 0)

	utils.DrawBadges(ctx, User, 10, 485)

	return ctx.Image()
}
