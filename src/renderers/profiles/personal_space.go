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

	ctx.DrawImage(userAvatar, 540, 75)

	backgroundImage := utils.GetAsset("/profiles/personal_space.png")

	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#845242")
	ctx.SetFontFace(*utils.GetFont("Mustard", 36))
	ctx.DrawStringWrapped("\u200e   "+User.Info, 827, 270, 0.5, 0, 455, 1, 0)

	ctx.DrawStringWrapped(User.Tag, 610, 105, 0, 0, 455, 1, 1)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Mustard", 32))
		ctx.DrawStringWrapped(User.Marry.Tag, 620, 140, 0, 0, 430, 1, 1)
	}

	utils.DrawVerticalBadges(ctx, User, 52, 38)

	ctx.SetFontFace(*utils.GetFont("Mustard", 32))
	ctx.DrawStringWrapped(I18n.Usages+". "+strconv.Itoa(int(User.Votes))+" upvotes", 610, 480, 0, 0, 460, 1, 0)

	return ctx.Image()
}
