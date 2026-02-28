package profiles

import (
	"image"
	"strconv"
	"strings"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/fogleman/gg"
)

func RenderNotes(User *utils.UserData, I18n *utils.I18n, db *database.Database) image.Image {
	ctx := gg.NewContext(1080, 720)

	backgroundImage := utils.GetAsset("/profiles/notes.png", db.ImageCache)
	userAvatar := utils.GetImageFromURL(User.Avatar, 236, 236, db)

	ctx.DrawImage(userAvatar, 18, 26)
	ctx.DrawImage(backgroundImage, 0, 0)

	utils.DrawBadges(ctx, db.ImageCache, User, 232, 636)

	ctx.SetFontFace(*utils.GetFont("Sans", 46))
	utils.StrokeText(ctx, User.Username, 280, 128, 2, 0, 0, "#000")
	ctx.SetHexColor("#FFF")
	ctx.DrawStringAnchored(User.Username, 280, 128, 0, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 34))
	ctx.SetHexColor("#000")
	ctx.DrawStringAnchored(User.Title, 270, 72, 0, 0)

	if User.Married {
		ctx.SetFontFace(*utils.GetFont("Sans", 32))
		ctx.SetHexColor("#FFF")
		ctx.DrawStringAnchored(User.MarryUsername, 280, 138, 0, 1)
	}

	ctx.SetFontFace(*utils.GetFont("Sans", 42))
	ctx.SetHexColor("#000")
	ctx.DrawStringWrapped(User.Info, 550, 480, 0.5, 1, 450, 1, gg.AlignCenter)

	User.Votes = 6666
	upvotesText := "Upvotes:" + strconv.Itoa(int(User.Votes))

	for i, char := range strings.Split(upvotesText, "") {
		ctx.DrawStringAnchored(char, float64(30+(27*i)), float64(639-(7*i)), 0, 0)
	}

	return ctx.Image()
}
