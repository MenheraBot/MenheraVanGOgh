package renderers

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type RouletteData struct {
	Bet    int64  `json:"bet"`    // 10 - 50000
	Drawn  int64  `json:"drawn"`  // 0 - 36
	Color  string `json:"color"`  // red green black
	Parity string `json:"parity"` // odd even
	Size   string `json:"size"`   // low high
	Dozen  string `json:"dozen"`  // first second third
}

func RenderRoulette(data *RouletteData, cache *database.Cache) image.Image {
	ctx := gg.NewContext(683, 1024)

	rouletteImage := utils.GetAsset("roulette/layout.png", cache)
	ctx.DrawImage(rouletteImage, 0, 0)

	if data.Color != "red" {
		replaceImageColor := utils.GetAsset(fmt.Sprintf("roulette/%s.png", data.Color), cache)

		ctx.DrawImage(replaceImageColor, 103, 648)
		ctx.DrawImage(replaceImageColor, 348, 648)
	}

	return ctx.Image()
}
