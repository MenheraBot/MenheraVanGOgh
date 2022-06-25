package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderFortification(User *utils.UserData, I18n *utils.I18n, util *utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	userAvatar := util.GetImageFromURL(User.Avatar, 250)
	ctx.DrawCircle(200, 200, 125)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 200, 200, 0.5, 0.5)
	ctx.ResetClip()

	ctx.SetHexColor(util.GetCompatibleFontColor(baseColor))
	ctx.SetFontFace(*util.GetFont("Sans", 45))
	ctx.DrawStringAnchored(I18n.Aboutme+":", 330, 310, 0, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 32))
	ctx.DrawStringWrapped(User.Info, 50, 330, 0, 0, 800, 1, 0)

	darkerColor := util.ShadeColor(baseColor, -15)
	ctx.SetHexColor(darkerColor)
	ctx.DrawRectangle(0, 480, 1080, 720)
	ctx.Fill()

	backgroundImage := util.GetAsset("/profiles/fortification.png")
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor(util.GetCompatibleFontColor(darkerColor))
	ctx.SetFontFace(*util.GetFont("Sans", 44))

	ctx.DrawStringWrapped(I18n.Usages+" | "+strconv.Itoa(int(User.Votes))+" Upvotes", 50, 545, 0, 0.5, 970, 1, 0)

	ctx.SetFontFace(*util.GetFont("Sans", 36))
	ctx.DrawStringAnchored(User.Tag, 630, 170, 0.5, 0)

	if User.Married {
		ctx.SetFontFace(*util.GetFont("Sans", 28))
		ctx.DrawStringWrapped(User.Marry.Username+" | "+strings.Split(User.MarryDate, " ")[0], 400, 240, 0, 0.5, 600, 1, 0)
		ringEmoji, _ := util.GetResizedAsset("badges/17.png", 42, 42)
		ctx.DrawImage(ringEmoji, 360, 220)
	}

	darkestThanTheDarkerColor := util.ShadeColor(darkerColor, -10)
	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(830, 270, 210, 200, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	ctx.SetFontFace(*util.GetFont("Sans", 40))
	ctx.SetHexColor(util.GetCompatibleFontColor(darkestThanTheDarkerColor))
	ctx.DrawStringAnchored(I18n.Mamado, 935, 310, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamou, 935, 400, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamadas)), 935, 355, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamou)), 935, 445, 0.5, 0)

	util.DrawBadges(ctx, User, 160, 5)

	return ctx.Image()
}
