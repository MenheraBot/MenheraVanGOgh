package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderPersonalSpace(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	userAvatar := utils.GetImageFromURL(User.Avatar, 250, db)
	backgroundImage := utils.GetAsset("/profiles/personal_space.png")

	ctx.DrawImage(userAvatar, 0, 0)
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#FFF")

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Sans", 42))
		ctx.DrawStringWrapped(User.Tag, 260, 90, 0, 0.5, 750, 1, 0)
		ctx.SetFontFace(*utils.GetFont("Sans", 38))
		ctx.DrawStringWrapped(User.Marry.Tag, 260, 150, 0, 0.5, 750, 1, 0)
	} else {
		ctx.SetFontFace(*utils.GetFont("Sans", 50))
		ctx.DrawStringWrapped(User.Tag, 260, 120, 0, 0.5, 750, 1, 0)
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 38))
	ctx.DrawStringWrapped(User.Info, 620, 220, 0.5, 0, 780, 1, 1)

	ctx.SetFontFace(*utils.GetFont("Sans", 34))
	ctx.DrawStringWrapped(I18n.Usages+"\n"+strconv.Itoa(int(User.Votes))+" Upvotes || "+I18n.Mamado+" "+strconv.Itoa(int(User.Mamadas))+" || "+I18n.Mamou+" "+strconv.Itoa(int(User.Mamou)), 550, 480, 0.5, 0, 920, 1, 1)

	utils.DrawBadges(ctx, User, 125, 637)

	return ctx.Image()
}
