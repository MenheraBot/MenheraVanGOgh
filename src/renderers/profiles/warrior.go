package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderWarrior(User *utils.UserData, I18n *utils.I18n, customEdits []string, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	if utils.GetProfileCustomization("useImage", customEdits) {
		backgroundImage := utils.GetImageFromURL(User.Image, 1080, 720, db)

		ctx.DrawImage(backgroundImage, 0, 0)
	} else {
		ctx.SetHexColor(baseColor)
		ctx.DrawRectangle(0, 0, 1080, 720)
		ctx.Fill()
	}

	userAvatar := utils.GetImageFromURL(User.Avatar, 226, 226, db)
	ctx.DrawImage(userAvatar, 23, 0)

	ctx.SetHexColor("#FFF")

	ctx.SetFontFace(*utils.GetFont("Sans", 36))
	ctx.DrawStringWrapped(User.Info, 105, 460, 0, 0.5, 870, 1, 0)

	ctx.DrawStringWrapped(I18n.Usages+"   | "+strconv.Itoa(int(User.Votes))+" Upvotes", 50, 275, 0, 0.5, 970, 1, 0)

	background := utils.GetAsset("profiles/guerreiro.png")
	ctx.DrawImage(background, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Warrior", 28))
	ctx.DrawStringAnchored(User.Tag, 330, 140, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Warrior", 16))
	ctx.DrawStringWrapped(User.Marry.Username+" "+strings.Split(User.MarryDate, " ")[0], 380, 170, 0, 0.5, 600, 1, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 28))
	ctx.DrawStringWrapped(I18n.Mamado+"\n"+strconv.Itoa(int(User.Mamou)), 940, 100, 0.5, 0.5, 600, 1, 1)
	ctx.DrawStringWrapped(I18n.Mamou+"\n"+strconv.Itoa(int(User.Mamadas)), 940, 170, 0.5, 0.5, 600, 1, 1)

	utils.DrawBadges(ctx, User, 110, 620)

	return ctx.Image()
}
