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

func RenderEightball(data *EightballData) image.Image {
	ctx := gg.NewContext(854, 456)

	bedroomImage := utils.GetAsset("images/bedroom.png")
	textBoxImage := utils.GetAsset("images/text_box.png")
	responseBoxImage, _ := utils.GetResizedAsset("images/response_box.png", 400, 250)
	menheraImage, _ := utils.GetResizedAsset("menheras/"+data.Type+"_"+strconv.Itoa(getRandomBasedOnType(data.Type))+".png", 387, 440)

	ctx.DrawImage(bedroomImage, 0, 0)
	ctx.DrawImage(menheraImage, 10, 10)
	ctx.DrawImage(textBoxImage, 40, 250)
	ctx.DrawImage(responseBoxImage, 440, 20)

	ctx.SetHexColor("#FFF")

	fontSize := 22
	if len(data.Username) <= 20 {
		fontSize = 28
	}

	ctx.SetFontFace(*utils.GetFont("Sans", float64(fontSize)))
	ctx.SetHexColor("#d89a30")
	ctx.DrawStringAnchored(data.Username, 440, 339, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 36))

	question := data.Question

	if !strings.HasSuffix(data.Question, "?") {
		question = data.Question + "?"
	}

	ctx.SetHexColor("#FFF")
	ctx.DrawStringWrapped(question, 440, 380, 0.5, 0.5, 700, 1, 1)

	ctx.SetHexColor("#595959")
	ctx.DrawStringWrapped(data.Answer, 645, 140, 0.5, 0.5, 360, 1, 1)

	return ctx.Image()
}
