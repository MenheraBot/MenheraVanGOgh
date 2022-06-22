package renderers

import (
	"image"
	"math/rand"
	"strconv"
	"strings"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type EightballData struct {
	Question string `json:"question"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Answer   string `json:"answer"`
}

const (
	Positive = 3
	Neutral  = 3
	Negative = 4
)

func getRandomBasedOnType(t string) int {
	switch t {
	case "positive":
		return rand.Intn(Positive) + 1
	case "neutral":
		return rand.Intn(Neutral) + 1
	case "negative":
		return rand.Intn(Negative) + 1
	default:
		return 1
	}
}

func RenderEightball(data *EightballData, util utils.Utils) image.Image {
	ctx := gg.NewContext(854, 456)

	bedroomImage := util.GetAsset("images/bedroom.png")
	textBoxImage := util.GetAsset("images/text_box.png")
	responseBoxImage, _ := util.GetResizedAsset("images/response_box.png", 400, 250)
	menheraImage, _ := util.GetResizedAsset("menheras/"+data.Type+"_"+strconv.Itoa(getRandomBasedOnType(data.Type))+".png", 387, 440)

	ctx.DrawImage(bedroomImage, 0, 0)
	ctx.DrawImage(menheraImage, 10, 10)
	ctx.DrawImage(textBoxImage, 40, 250)
	ctx.DrawImage(responseBoxImage, 440, 20)

	ctx.SetHexColor("#FFF")

	fontSize := 22
	if len(data.Username) <= 20 {
		fontSize = 28
	}

	ctx.LoadFontFace(util.GetFontPath("Arial"), float64(fontSize))
	util.StrokeText(ctx, data.Username, 440, 339, 2, "#000", "#FFF", 0.5)

	ctx.LoadFontFace(util.GetFontPath("Arial"), 38)

	question := data.Question

	if !strings.HasSuffix(data.Question, "?") {
		question = data.Question + "?"
	}

	util.FillStrokedText(ctx, question, 440, 378, 700, 500, 30, 2, "#000", "#FFF", 0.5)

	util.FillStrokedText(ctx, data.Answer, 645, 80, 360, 500, 35, 2, "#000", "#FFF", 0.5)

	return ctx.Image()
}
