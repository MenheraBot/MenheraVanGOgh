package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderID03(User *utils.UserData, I18n *utils.I18n) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	userAvatar := utils.GetImageFromURL(User.Avatar, 120)

	ctx.DrawImage(userAvatar, 225, 92)

	backgroundImage := utils.GetAsset("/profiles/iD03.png")

	ctx.DrawImage(backgroundImage, 0, 0)

	utils.DrawBadges(ctx, User, 408, 435)

	ctx.SetHexColor("#FFF")
	ctx.SetFontFace(*utils.GetFont("Pixellari", 32))
	ctx.DrawStringWrapped(User.Tag, 425, 160, 0, 1, 420, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Pixellari", 24))
	ctx.DrawStringWrapped(User.Info, 200, 540, 0, 0.5, 700, 1, 0)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Pixellari", 20))
		ctx.DrawStringWrapped(User.Marry.Username+" "+strings.Split(User.MarryDate, " ")[0], 445, 220, 0, 1, 600, 1, 0)
	}

	fontSize := 24

	if User.Mamou >= 1000 || User.Mamadas >= 100 {
		fontSize = 22
	}

	ctx.SetFontFace(*utils.GetFont("Pixellari", float64(fontSize)))
	ctx.DrawStringAnchored(I18n.Mamado+": "+strconv.Itoa(int(User.Mamadas)), 850, 145, 0, 0)
	ctx.DrawStringAnchored(I18n.Mamou+": "+strconv.Itoa(int(User.Mamou)), 850, 175, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Pixellari", 24))
	ctx.DrawStringWrapped(strings.Split(I18n.Usages, ".")[0]+" | "+strconv.Itoa(int(User.Votes))+" Upvotes", 390, 426, 0, 1, 540, 1, 0)

	return ctx.Image()
}
