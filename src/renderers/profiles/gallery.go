package profiles

import (
	"image"
	"strconv"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderGallery(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	ctx.SetHexColor(User.Color)
	ctx.DrawRectangle(0, 0, 1080, 720)
	ctx.Fill()

	userAvatar := utils.GetImageFromURL(User.Avatar, 220, 220, db)
	ctx.DrawCircle(170, 160, 110)
	ctx.Clip()
	ctx.DrawImageAnchored(userAvatar, 170, 160, 0.5, 0.5)
	ctx.ResetClip()

	backgroundImage := utils.GetAsset("/profiles/gallery.png")
	ctx.DrawImage(backgroundImage, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Arial", 36))

	ctx.SetHexColor("#000")
	ctx.DrawStringAnchored(User.Title, 687, 650, 0.5, 0.5)

	ctx.SetHexColor(utils.GetCompatibleFontColor(User.Color))

	ctx.SetFontFace(*utils.GetFont("Sans", 42))
	ctx.DrawStringWrapped(User.Username, 540, 110, 0.5, 1, 350, 1, gg.AlignCenter)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Sans", 38))
		ctx.DrawStringWrapped(User.MarryUsername, 540, 120, 0.5, 0, 350, 1, gg.AlignCenter)
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 36))
	ctx.DrawStringWrapped(User.Info, 680, 465, 0.5, 0.5, 420, 1, gg.AlignCenter)

	ctx.SetFontFace(*utils.GetFont("Sans", 28))
	ctx.DrawStringAnchored("Upvotes "+strconv.Itoa(int(User.Votes)), 940, 100, 0.5, 0.5)
	ctx.DrawStringAnchored(I18n.Mamou+" "+strconv.Itoa(int(User.Mamou)), 940, 140, 0.5, 0.5)
	ctx.DrawStringAnchored(I18n.Mamado+" "+strconv.Itoa(int(User.Mamadas)), 940, 178, 0.5, 0.5)

	utils.DrawBadgesWrapped(ctx, db.ImageCache, User, 100, 366, 3)

	return ctx.Image()
}
