package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderSunflower(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	ctx.SetHexColor(User.Color)
	ctx.DrawRectangle(186, 38, 766, 181)
	ctx.Fill()

	userAvatar := utils.GetImageFromURL(User.Avatar, 140, 140, db)
	ctx.DrawCircle(136, 126, 70)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 136, 126, 0.5, 0.5)
	ctx.ResetClip()

	backgroundImage := utils.GetAsset("/profiles/sunflower.png", db.ImageCache)
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetHexColor(utils.GetCompatibleFontColor(User.Color))

	ctx.SetFontFace(*utils.GetFont("Sans", 46))
	ctx.DrawString(User.Username, 270, 90)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Arial", 40))
		ctx.DrawString(User.MarryUsername, 270, 140)
	}

	ctx.SetHexColor("#000")
	ctx.SetFontFace(*utils.GetFont("Arial", 40))
	ctx.DrawStringWrapped(User.Info, 595, 230, 0.5, 0, 680, 1, gg.AlignCenter)

	ctx.SetFontFace(*utils.GetFont("Sans", 36))
	ctx.DrawStringAnchored("Upvotes", 710, 480, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamou, 710, 520, 0.5, 0)
	ctx.DrawStringAnchored(I18n.Mamado, 710, 560, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("Arial", 36))
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Votes)), 820, 480, 0, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamou)), 820, 520, 0, 0)
	ctx.DrawStringAnchored(strconv.Itoa(int(User.Mamadas)), 820, 560, 0, 0)

	ctx.SetHexColor("#fff")
	ctx.SetFontFace(*utils.GetFont("Arial", 38))
	ctx.DrawStringAnchored(User.Title, 300, 660, 0.5, 0)

	utils.DrawBadges(ctx, db.ImageCache, User, 270, 155)

	return ctx.Image()
}
