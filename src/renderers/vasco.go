package renderers

import (
	"image"
	"strings"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type VascoData struct {
	User     string `json:"user"`
	Username string `json:"username"`
	Quality  string `json:"quality"`
	Position string `json:"position"`
}

func RenderVasco(data *VascoData, util utils.Utils) image.Image {
	ctx := gg.NewContext(800, 534)

	vascoImage := util.GetAsset("images/vasco_" + data.Quality + ".png")
	userImage := util.GetImageFromURL(data.User, 243, 243)

	ctx.DrawImage(userImage, 65, 165)
	ctx.DrawImage(vascoImage, 0, 0)
	ctx.LoadFontFace(util.GetFontPath("Impact"), 42)

	ctx.SetHexColor("#FFF")

	text := strings.ToUpper(data.Username) + "\n" + strings.ToUpper(data.Position)

	util.FillStrokedText(ctx, text, 500, 260, 350, 500, 50, 0, "#FFF", "#FFF", 0.5)

	return ctx.Image()
}
