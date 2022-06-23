package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderKawaii(User *utils.UserData, I18n *utils.I18n, util utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	userAvatar := util.GetImageFromURL(User.Avatar, 300)

	ctx.DrawImage(userAvatar, 55, 50)

	backgroundImage := util.GetAsset("/profiles/kawaii.png")

	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#FFF")
	ctx.SetFontFace(*util.GetFont("Kawaii", 60))

	ctx.DrawStringAnchored(I18n.Mamou, 880, 440, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamado, 880, 560, 0.5, 0)

	ctx.DrawStringAnchored(strconv.Itoa(User.Mamou), 880, 500, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(User.Mamadas), 880, 620, 0.5, 0)

	ctx.SetFontFace(*util.GetFont("Kawaii", 72))
	ctx.DrawStringAnchored(User.Tag, 420, 200, 0, 0)

	if User.Married {
		ctx.SetFontFace(*util.GetFont("Kawaii", 36))
		ctx.DrawStringWrapped(User.Marry.Username+" "+strings.Split(User.MarryDate, " ")[0], 460, 290, 0, 1, 600, 1, 0)
		ringEmoji, _ := util.GetResizedAsset("/badges/17.png", 42, 42)
		ctx.DrawImage(ringEmoji, 415, 260)
	}

	ctx.SetFontFace(*util.GetFont("Kawaii", 32))
	ctx.DrawStringWrapped(User.Info, 85, 410, 0, 0.5, 680, 1, 0)

	ctx.SetFontFace(*util.GetFont("Kawaii", 34))
	ctx.DrawStringWrapped(I18n.Usages+"   | "+strconv.Itoa(User.Votes)+" Upvotes", 85, 580, 0, 0.5, 650, 1, 0)

	util.DrawBadges(ctx, User, 410, 40)
	/*
		util.DrawBadges(ctx, User, 408, 435)

		ctx.DrawStringWrapped(User.Tag, 425, 160, 0, 1, 420, 0, 0)

		ctx.SetFontFace(*util.GetFont("Pixellari", 24))
		ctx.DrawStringWrapped(User.Info, 200, 540, 0, 0.5, 700, 1, 0)

		if User.Married {
			ctx.SetFontFace(*util.GetFont("Pixellari", 20))
			ctx.DrawStringWrapped(User.Marry.Username+" "+strings.Split(User.MarryDate, " ")[0], 445, 220, 0, 1, 600, 1, 0)
		}

		fontSize := 24

		if User.Mamou >= 1000 || User.Mamadas >= 100 {
			fontSize = 22
		}

		ctx.SetFontFace(*util.GetFont("Pixellari", float64(fontSize)))
		ctx.DrawStringAnchored(I18n.Mamado+": "+strconv.Itoa(User.Mamadas), 850, 145, 0, 0)
		ctx.DrawStringAnchored(I18n.Mamou+": "+strconv.Itoa(User.Mamou), 850, 175, 0, 0)

		ctx.SetFontFace(*util.GetFont("Pixellari", 24))
		ctx.DrawStringWrapped(strings.Split(I18n.Usages, ".")[0]+" | "+strconv.Itoa(User.Votes)+" Upvotes", 390, 426, 0, 1, 540, 1, 0)
	*/
	return ctx.Image()
}
