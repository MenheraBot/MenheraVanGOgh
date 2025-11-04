package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderMemories(User *utils.UserData, I18n *utils.I18n, customEdits []string, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	if utils.GetProfileCustomization("useImage", customEdits) {
		backgroundImage := utils.GetImageFromURL(User.Image, 1080, 720, db)
		ctx.DrawImage(backgroundImage, 110, 90)
	} else {
		ctx.SetHexColor(User.Color)
		ctx.DrawRectangle(110, 90, 950, 640)
		ctx.Fill()
	}

	userAvatar := utils.GetImageFromURL(User.Avatar, 226, 226, db)
	ctx.DrawCircle(180, 160, 120)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 170, 155, 0.5, 0.5)
	ctx.ResetClip()

	backgroundImage := utils.GetAsset("/profiles/memories.png")
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 56))
	ctx.SetHexColor("#fff")
	ctx.DrawString(User.Username, 380, 70)

	upperTextStrokeColor := "#000"
	upperTextColor := "#FFF"

	if utils.GetProfileCustomization("whiteUpperText", customEdits) {
		ctx.SetHexColor(upperTextColor)
	} else {
		upperTextStrokeColor = "#FFF"
		upperTextColor = "#000"

		ctx.SetHexColor(upperTextColor)
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 32))

	ctx.DrawStringAnchored(strings.ToUpper(I18n.Mamado), 478, 120, 0.5, 0.5)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamadas)), 478, 160, 0.5, 0.5)

	ctx.DrawStringAnchored(strings.ToUpper(I18n.Mamou), 478, 200, 0.5, 0.5)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamou)), 478, 240, 0.5, 0.5)

	ctx.SetFontFace(*utils.GetFont("Sans", 26))

	utils.StrokeText(ctx, User.Title, 855, 183, 2, 0.5, 0.5, upperTextStrokeColor)
	ctx.SetHexColor(upperTextColor)
	ctx.DrawStringAnchored(User.Title, 855, 183, 0.5, 0.5)

	utils.StrokeText(ctx, strconv.Itoa(int(User.Votes))+" Upvotes", 855, 217, 2, 0.5, 0.5, upperTextStrokeColor)
	ctx.SetHexColor(upperTextColor)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Votes))+" Upvotes", 855, 217, 0.5, 0.5)

	if User.Married {
		utils.StrokeText(ctx, User.MarryUsername, 855, 145, 2, 0.5, 0.5, upperTextStrokeColor)

		ctx.SetHexColor(upperTextColor)
		ctx.DrawStringAnchored(User.MarryUsername, 855, 145, 0.5, 0.5)
	}

	if utils.GetProfileCustomization("whiteBottomText", customEdits) {
		ctx.SetHexColor("#FFF")
	} else {
		ctx.SetHexColor("#000")
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 42))
	ctx.DrawString(I18n.Aboutme, 170, 408)

	ctx.SetFontFace(*utils.GetFont("Sans", 40))
	ctx.DrawStringWrapped(User.Info, 120, 420, 0, 0, 900, 1, gg.AlignLeft)

	splittedUsages := strings.Split(I18n.Usages, ".")

	ctx.SetFontFace(*utils.GetFont("Sans", 34))
	ctx.DrawStringWrapped(splittedUsages[0], 160, 580, 0, 0, 400, 1, gg.AlignLeft)

	if len(splittedUsages) > 1 {
		ctx.DrawStringWrapped(splittedUsages[1], 600, 580, 0, 0, 430, 1, gg.AlignLeft)
	}

	utils.DrawBadges(ctx, db.ImageCache, User, 120, 271)

	return ctx.Image()
}
