package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderWithoutSoul(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Color

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	userAvatar := utils.GetImageFromURL(User.Avatar, 175, 175, db)

	ctx.DrawImage(userAvatar, 85, 95)

	backgroundImage := utils.GetAsset("profiles/without_soul.png")

	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor("#FFF")
	ctx.SetFontFace(*utils.GetFont("Postamt", 22))
	ctx.DrawStringWrapped(User.Info, 385, 200, 0, 0.5, 520, 1.3, 0)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Gabrielle", 32))
		ctx.DrawStringWrapped(User.MarryUsername+" | "+strings.Split(User.MarryDate, " ")[0], 440, 325, 0, 1, 600, 1, 0)
	}

	ctx.SetFontFace(*utils.GetFont("Postamt", 30))

	ctx.DrawStringAnchored(User.Title, 630, 73, 0.5, 0.5)

	if len(User.Username) > 22 {
		ctx.SetFontFace(*utils.GetFont("Postamt", 24))
	}

	ctx.DrawStringWrapped(User.Username, 630, 125, 0.5, 0.5, 420, 1, 1)

	ctx.DrawStringAnchored(User.Title, 630, 73, 0.5, 0.5)

	ctx.SetFontFace(*utils.GetFont("Postamt", 30))
	ctx.DrawStringAnchored("Upvotes: "+strconv.Itoa(int(User.Votes)), 875, 620, 0.5, 0.5)

	utils.DrawBadges(ctx, db.ImageCache, User, 135, 590)

	return ctx.Image()
}
