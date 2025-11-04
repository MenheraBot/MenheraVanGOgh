package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderWebsite(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	backgroundImage := utils.GetAsset("/profiles/website.png")
	userAvatar := utils.GetImageFromURL(User.Avatar, 226, 226, db)

	ctx.DrawImage(userAvatar, 23, 140)
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#000")
	ctx.SetFontFace(*utils.GetFont("Arial", 30))
	ctx.DrawString(User.Username, 300, 100)

	if User.Married {
		marryName := ctx.WordWrap(User.MarryUsername, 300)
		marryfirstName := marryName[0]
		if len(marryName) > 1 {
			marryfirstName += "..."
		}

		ctx.DrawStringAnchored(marryfirstName, 135, 440, 0.5, 0)
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 38))

	userName := ctx.WordWrap(User.Username, 300)
	userFirstName := userName[0]
	if len(userName) > 1 {
		userFirstName += "..."
	}

	ctx.DrawStringAnchored(userFirstName, 135, 400, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 48))
	ctx.DrawString(I18n.Aboutme, 360, 170)

	ctx.SetFontFace(*utils.GetFont("Sans", 36))

	ctx.DrawStringAnchored("Upvotes "+strconv.Itoa(int(User.Votes)), 135, 500, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamado+" "+strconv.Itoa(int(User.Mamadas)), 135, 550, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamou+" "+strconv.Itoa(int(User.Mamou)), 135, 600, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("Arial", 46))
	ctx.DrawStringWrapped(User.Info, 700, 200, 0.5, 0, 700, 1.2, gg.AlignCenter)

	ctx.SetHexColor("#fff")
	ctx.SetFontFace(*utils.GetFont("Postamt", 38))
	ctx.DrawStringAnchored(User.Title, 810, 613, 0.5, 0)

	utils.DrawBadges(ctx, db.ImageCache, User, 0, 653)

	return ctx.Image()
}
