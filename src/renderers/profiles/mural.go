package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderPersonalSpace(User *utils.UserData, I18n *utils.I18n, customEdits []string, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	userAvatar := utils.GetImageFromURL(User.Avatar, 250, 250, db)
	backgroundImage := utils.GetImageFromURL(User.Image, 1080, 720, db)

	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor(User.Color)
	ctx.DrawCircle(131, 131, 130)
	ctx.Fill()
	ctx.DrawCircle(131, 131, 110)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 131, 131, 0.5, 0.5)
	ctx.ResetClip()

	ctx.MoveTo(235, 60)
	ctx.LineTo(892, 60)
	ctx.LineTo(1030, 130)
	ctx.LineTo(892, 202)
	ctx.LineTo(235, 202)

	ctx.SetLineWidth(15)

	if utils.GetProfileCustomization("textBoxFilled", customEdits) {
		ctx.Fill()
	} else {
		ctx.Stroke()
	}

	ctx.MoveTo(47, 465)
	ctx.LineTo(1030, 465)
	ctx.LineTo(958, 574)
	ctx.LineTo(1030, 682)
	ctx.LineTo(47, 682)
	ctx.LineTo(121, 574)
	ctx.LineTo(47, 465)

	if utils.GetProfileCustomization("textBoxFilled", customEdits) {
		ctx.Fill()
	} else {
		ctx.Stroke()
	}

	ctx.SetHexColor(utils.ShadeColor(User.Color, 15))

	ctx.DrawRectangle(128, 635, 822, 66)
	ctx.Fill()

	if utils.GetProfileCustomization("forceWhiteInfo", customEdits) {
		ctx.SetHexColor("#FFF")
	} else {
		ctx.SetHexColor(utils.GetCompatibleFontColor(User.Color))
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 38))
	ctx.DrawStringWrapped(User.Info, 620, 220, 0.5, 0, 780, 1, 1)

	ctx.SetHexColor(utils.GetCompatibleFontColor(User.Color))

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Sans", 42))
		ctx.DrawStringWrapped(User.Tag, 260, 90, 0, 0.5, 750, 1, 0)
		ctx.SetFontFace(*utils.GetFont("Sans", 38))
		ctx.DrawStringWrapped(User.Marry.Tag, 260, 150, 0, 0.5, 750, 1, 0)
	} else {
		ctx.SetFontFace(*utils.GetFont("Sans", 50))
		ctx.DrawStringWrapped(User.Tag, 260, 120, 0, 0.5, 750, 1, 0)
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 34))
	ctx.DrawStringWrapped(I18n.Usages+"\n"+strconv.Itoa(int(User.Votes))+" Upvotes || "+I18n.Mamado+" "+strconv.Itoa(int(User.Mamadas))+" || "+I18n.Mamou+" "+strconv.Itoa(int(User.Mamou)), 550, 480, 0.5, 0, 920, 1, 1)

	utils.DrawBadges(ctx, User, 125, 637)

	return ctx.Image()
}
