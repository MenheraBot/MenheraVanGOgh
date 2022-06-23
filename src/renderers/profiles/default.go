package profiles

import (
	"image"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderDefault(User *utils.UserData, I18n *utils.I18n, util utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Cor

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	darker := util.ShadeColor(baseColor, -15)

	ctx.SetHexColor(darker)
	ctx.DrawRoundedRectangle(0, 0, 1080, 240, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	darkestThanTheDarkerColor := util.ShadeColor(darker, -10)

	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(0, 164, 1080, 75, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(890, 250, 180, 200, 20)
	ctx.FillPreserve()
	ctx.SetHexColor("#000")
	ctx.Stroke()

	userAvatar := util.GetImageFromURL(User.Avatar, 250, 250)

	ctx.SetHexColor("#000")
	ctx.DrawCircle(120, 120, 130)
	ctx.Fill()

	ctx.DrawCircle(120, 120, 120)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 120, 120, 0.5, 0.5)
	ctx.ResetClip()

	ctx.SetHexColor("#FFF")

	/* ctx.SetFontFace(*util.GetFont("Sans", 50))
	util.FillStrokedText(ctx, User.Tag, 255, 90, 650, 300, 50, 3, "#000", "#FFF", 0)

	ctx.SetFontFace(*util.GetFont("Sans", 45))
	util.StrokeText(ctx, "Upvotes", 860, 60, 3, "#000", "#FFF", 0)
	util.StrokeText(ctx, strconv.Itoa(User.Votos), 955, 120, 2, "#000", "#FFF", 0.5)

	ctx.SetFontFace(*util.GetFont("Sans", 55))
	util.StrokeText(ctx, I18n.Aboutme, 20, 300, 3, "#000", "#FFF", 0)

	ctx.SetFontFace(*util.GetFont("Sans", 40))
	util.FillStrokedText(ctx, User.Nota, 20, 350, 870, 600, 40, 2, "#000", "#FFF", 0)

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 480, 1080, 720)
	ctx.Fill()

	ctx.SetFontFace(*util.GetFont("Sans", 50))
	util.FillStrokedText(ctx, I18n.Usages, 20, 600, 1000, 600, 50, 2, "#000", "#FFF", 0)

	ctx.SetFontFace(*util.GetFont("Sans", 40))
	if User.Married {
		ringEmoji, _ := util.GetResizedAsset("badges/17.png", 64, 64)
		util.StrokeText(ctx, User.Marry.Tag+" | "+User.Data, 80, 535, 2, "#000", "#FFF", 0)
		ctx.DrawImage(ringEmoji, 10, 490)
	}

	util.StrokeText(ctx, I18n.Mamado, 980, 290, 3, "#000", "#FFF", 0.5)
	util.StrokeText(ctx, I18n.Mamou, 980, 380, 3, "#000", "#FFF", 0.5)

	util.StrokeText(ctx, strconv.Itoa(User.Mamadas), 980, 335, 2, "#000", "#FFF", 0.5)
	util.StrokeText(ctx, strconv.Itoa(User.Mamou), 980, 425, 2, "#000", "#FFF", 0.5)
	*/
	util.DrawBadges(ctx, User, 230, 170)

	return ctx.Image()
}
