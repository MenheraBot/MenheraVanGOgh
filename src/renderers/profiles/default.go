package profiles

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func roundImage(img image.Image) image.Image {
	ctx := gg.NewContext(250, 250)

	ctx.DrawRoundedRectangle(0, 0, 240, 240, 120)
	ctx.Clip()
	ctx.DrawImage(img, 0, 0)

	return ctx.Image()
}

func RenderDefault(User *utils.UserData, UsageCommands *utils.UsageCommands, I18n *utils.I18n, util utils.Utils) image.Image {
	ctx := gg.NewContext(1080, 720)

	baseColor := User.Cor

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	darker := util.ShadeColor(baseColor, -15)

	ctx.SetHexColor(darker)
	ctx.DrawRoundedRectangle(0, 0, 1080, 240, 20)
	ctx.Fill()

	darkestThanTheDarkerColor := util.ShadeColor(darker, -10)

	ctx.SetHexColor(darkestThanTheDarkerColor)
	ctx.DrawRoundedRectangle(0, 164, 1080, 75, 20)
	ctx.Fill()

	ctx.DrawRoundedRectangle(890, 250, 180, 200, 20)
	ctx.Fill()

	userAvatar := util.GetImageFromURL(User.Avatar, 250, 250)

	ctx.SetHexColor("#000")
	ctx.DrawCircle(120, 120, 130)
	ctx.Fill()

	ctx.DrawImage(roundImage(userAvatar), 0, 0)

	ctx.SetHexColor("#FFF")

	ctx.LoadFontFace(util.GetFontPath("Arial"), 50)
	util.FillStrokedText(ctx, User.Tag, 255, 90, 650, 300, 40, 2, "#000", "#FFF", 0)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 45)
	util.StrokeText(ctx, "Upvotes", 860, 60, 2, "#000", "#FFF", 0)
	util.StrokeText(ctx, strconv.Itoa(User.Votos), 955, 120, 2, "#000", "#FFF", 0.5)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 55)
	util.StrokeText(ctx, I18n.Aboutme, 20, 300, 3, "#000", "#FFF", 0)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 40)
	util.FillStrokedText(ctx, User.Nota, 20, 350, 870, 130, 30, 2, "#000", "#FFF", 0)

	ctx.SetHexColor(baseColor)
	ctx.DrawRectangle(0, 480, 1080, 720)

	if UsageCommands != nil {
		usedCommands := UsageCommands.Cmds.Count
		mostUsedCommand := UsageCommands.Array[0]

		text := fmt.Sprintf("%s %s %d %s %s %d %s", User.Username, I18n.Zero, usedCommands, strings.ToTitle(mostUsedCommand.Name), I18n.Dois, mostUsedCommand.Count, I18n.Tres)

		util.FillStrokedText(ctx, text, 20, 600, 1000, 600, 30, 2, "#000", "#FFF", 0)
	}

	if User.Married {
		ringEmoji, _ := util.GetResizedAsset("badges/17.png", 64, 64)
		ctx.SetLineWidth(1)
		util.StrokeText(ctx, User.Marry.Tag+" | "+User.Data, 80, 535, 2, "#000", "#FFF", 0)
		ctx.DrawImage(ringEmoji, 10, 490)
	}

	util.StrokeText(ctx, I18n.Mamado, 980, 290, 2, "#000", "#FFF", 0.5)
	util.StrokeText(ctx, I18n.Mamou, 980, 380, 2, "#000", "#FFF", 0.5)

	util.StrokeText(ctx, strconv.Itoa(User.Mamadas), 980, 335, 2, "#000", "#FFF", 0.5)
	util.StrokeText(ctx, strconv.Itoa(User.Mamou), 980, 425, 2, "#000", "#FFF", 0.5)

	util.DrawBadges(ctx, User, 230, 170)

	return ctx.Image()
}
