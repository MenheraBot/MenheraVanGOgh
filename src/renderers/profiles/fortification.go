package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderFortification(User *utils.UserData, I18n *utils.I18n, customEdits []string, db *database.Database) image.Image {
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

	userAvatar := utils.GetImageFromURL(User.Avatar, 250, 250, db)
	ctx.DrawCircle(200, 200, 125)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 200, 200, 0.5, 0.5)
	ctx.ResetClip()

	ctx.SetHexColor(utils.GetCompatibleFontColor(baseColor))
	ctx.SetFontFace(*utils.GetFont("Sans", 45))
	ctx.DrawStringAnchored(I18n.Aboutme+":", 330, 310, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 32))
	ctx.DrawStringWrapped(User.Info, 50, 330, 0, 0, 800, 1, 0)

	darkerColor := utils.ShadeColor(baseColor, -15)

	if !utils.GetProfileCustomization("useImage", customEdits) {
		ctx.SetHexColor(darkerColor)
		ctx.DrawRectangle(0, 480, 1080, 720)
		ctx.Fill()
	}

	backgroundImage := utils.GetAsset("/profiles/fortification.png", db.ImageCache)
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor(utils.GetCompatibleFontColor(darkerColor))
	ctx.SetFontFace(*utils.GetFont("Sans", 44))

	ctx.DrawStringWrapped(I18n.Usages+" | "+strconv.Itoa(int(User.Votes))+" Upvotes", 50, 545, 0, 0.5, 970, 1, 0)

	ctx.SetHexColor("#fff")
	ctx.SetFontFace(*utils.GetFont("Sans", 36))
	ctx.DrawStringAnchored(User.Username, 630, 140, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("Arial", 36))
	utils.StrokeText(ctx, User.Title, 630, 190, 2, 0.5, 0, "#000")
	ctx.SetHexColor("#fff")
	ctx.DrawStringAnchored(User.Title, 630, 190, 0.5, 0)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Sans", 28))
		ctx.DrawStringWrapped(User.MarryUsername+" | "+strings.Split(User.MarryDate, " ")[0], 400, 240, 0, 0.5, 600, 1, 0)
		ringEmoji, _ := utils.GetResizedAsset("badges/17.png", 42, 42)
		ctx.DrawImage(ringEmoji, 360, 220)
	}

	darkestThanTheDarkerColor := utils.ShadeColor(darkerColor, -10)

	if !utils.GetProfileCustomization("useImage", customEdits) {
		ctx.SetHexColor(darkestThanTheDarkerColor)
		ctx.DrawRoundedRectangle(830, 270, 210, 200, 20)
		ctx.FillPreserve()
		ctx.SetHexColor("#000")
		ctx.Stroke()
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 40))
	ctx.SetHexColor(utils.GetCompatibleFontColor(darkestThanTheDarkerColor))
	ctx.DrawStringAnchored(I18n.Mamado, 935, 310, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamou, 935, 400, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamadas)), 935, 355, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamou)), 935, 445, 0.5, 0)

	utils.DrawBadges(ctx, db.ImageCache, User, 160, 5)

	return ctx.Image()
}
