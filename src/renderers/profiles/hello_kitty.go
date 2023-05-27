package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderHelloKitty(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	userAvatar := utils.GetImageFromURL(User.Avatar, 200, db)
	backgroundImage := utils.GetAsset("/profiles/hello_kitty.png")

	ctx.DrawImage(userAvatar, 20, 0)
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#FFF")

	ctx.SetFontFace(*utils.GetFont("Kawaii", 54))
	ctx.DrawStringWrapped(User.Tag, 355, 100, 0, 0.5, 750, 1, 0)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Kawaii", 32))
		ctx.DrawStringWrapped(User.Marry.Tag, 355, 140, 0, 0.5, 750, 1, 0)
	}

	ctx.SetFontFace(*utils.GetFont("Kawaii", 38))
	ctx.DrawStringWrapped(User.Info, 625, 170, 0.5, 0, 560, 1, 0)

	ctx.SetFontFace(*utils.GetFont("Kawaii", 34))
	ctx.DrawStringWrapped(I18n.Usages+"\n"+strconv.Itoa(int(User.Votes))+" Upvotes || "+I18n.Mamado+" "+strconv.Itoa(int(User.Mamadas))+" || "+I18n.Mamou+" "+strconv.Itoa(int(User.Mamou)), 320, 410, 0.5, 0, 500, 1, 1)

	utils.DrawVerticalBadges(ctx, User, 952, 70)

	return ctx.Image()
}
