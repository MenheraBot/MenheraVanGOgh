package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderFrozen(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	userAvatar := utils.GetImageFromURL(User.Avatar, 270, 260, db)
	backgroundImage := utils.GetAsset("/profiles/frozen.png")
	ctx.DrawImage(userAvatar, 34, 23)
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#000")
	ctx.SetFontFace(*utils.GetFont("IceCaps", 60))
	ctx.DrawStringAnchored(User.Title, 670, 40, 0.5, 0.5)

	ctx.DrawStringAnchored(I18n.Aboutme, 560, 470, 0.5, 0.5)

	ctx.SetFontFace(*utils.GetFont("Icetea", 48))
	ctx.DrawStringWrapped(User.Info, 550, 575, 0.5, 0.5, 900, 1, gg.AlignCenter)

	ctx.SetFontFace(*utils.GetFont("Icetea", 36))
	ctx.DrawStringAnchored(I18n.Mamado, 455, 237, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamadas)), 455, 265, 0.5, 0)

	ctx.DrawStringAnchored(I18n.Mamou, 455, 290, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamou)), 455, 317, 0.5, 0)

	ctx.DrawStringAnchored("Upvotes", 695, 250, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Votes)), 695, 300, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("IceCaps", 48))
	ctx.DrawStringAnchored(User.Username, 165, 330, 0.5, 0)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("IceCaps", 34))
		ctx.DrawStringAnchored(User.MarryUsername, 175, 380, 0.5, 0)
		ctx.SetFontFace(*utils.GetFont("Icetea", 32))
		ctx.DrawStringAnchored(User.MarryDate, 175, 410, 0.5, 0)
	}

	utils.DrawBadges(ctx, db.ImageCache, User, 370, 350)

	return ctx.Image()
}
