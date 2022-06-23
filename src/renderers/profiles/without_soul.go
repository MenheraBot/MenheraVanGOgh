package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderWithoutSoul(User *utils.UserData, I18n *utils.I18n, util utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	userAvatar := util.GetImageFromURL(User.Avatar, 175)

	ctx.DrawImage(userAvatar, 85, 95)

	backgroundImage := util.GetAsset("profiles/without_soul.png")

	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#FFF")
	ctx.SetFontFace(*util.GetFont("Postamt", 22))
	ctx.DrawStringWrapped(User.Info, 385, 200, 0, 0.5, 520, 1.3, 0)

	if User.Married {
		ctx.SetFontFace(*util.GetFont("Gabrielle", 32))
		ctx.DrawStringWrapped(User.Marry.Username+" | "+strings.Split(User.MarryDate, " ")[0], 440, 325, 0, 1, 600, 1, 0)
	}

	fontSize := 30

	if len(User.Username) > 22 {
		fontSize = 24
	}

	ctx.SetFontFace(*util.GetFont("Postamt", float64(fontSize)))
	ctx.DrawStringWrapped(User.Tag, 630, 127, 0.5, 0.5, 420, 1, 1)

	ctx.SetFontFace(*util.GetFont("Postamt", 30))
	ctx.DrawStringAnchored("Upvotes: "+strconv.Itoa(User.Votes), 875, 620, 0.5, 0.5)

	util.DrawBadges(ctx, User, 135, 595)

	return ctx.Image()
}
