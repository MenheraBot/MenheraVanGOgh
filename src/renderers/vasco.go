package renderers

import (
	"image"
	"strings"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type VascoData struct {
	User     string `json:"user"` // 256 normal | low 64
	Username string `json:"username"`
	Quality  string `json:"quality"`
	Position string `json:"position"`
}

func RenderVasco(data *VascoData, db *database.Database) image.Image {
	ctx := gg.NewContext(800, 534)

	vascoImage := utils.GetAsset("images/vasco_" + data.Quality + ".png", db.ImageCache)
	userImage := utils.GetImageFromURL(data.User, 243, 243, db)

	ctx.DrawImage(userImage, 65, 165)
	ctx.DrawImage(vascoImage, 0, 0)
	ctx.SetFontFace(*utils.GetFont("Impact", 42))

	text := strings.ToUpper(data.Username) + "\n" + strings.ToUpper(data.Position)

	ctx.SetHexColor("#FFF")
	ctx.DrawStringWrapped(text, 500, 290, 0.5, 0.5, 500, 2, 1)

	return ctx.Image()
}
