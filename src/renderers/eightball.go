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
	Question        string `json:"question"`
	Type            string `json:"type"`
	Username        string `json:"username"`
	Answer          string `json:"answer"`
	BackgroundTheme string `json:"backgroundTheme"`
	TextBoxTheme    string `json:"textBoxTheme"`
	MenheraTheme    string `json:"menheraTheme"`
}

func getRandomMenheraImage(t, theme string) string {
	const MenherasByType = 5

	return "8ball/menheras/" + theme + "/" + t + "_" + strconv.Itoa(rand.Intn(MenherasByType)+1) + ".png"
}

func getQuestionTextColor(theme string) string {
	switch theme {
	case "default":
		return "#FFF"
	case "xp":
		return "#000"
	default:
		return "#FFF"
	}
}

func getResponseTextColor(theme string) string {
	switch theme {
	case "default":
		return "#595959"
	case "xp":
		return "#000"
	default:
		return "#595959"
	}
}

func getUsernameTextColor(theme string) string {
	switch theme {
	case "default":
		return "#d89a30"
	case "xp":
		return "#FFF"
	default:
		return "#d89a30"
	}
}

func RenderEightball(data *EightballData) image.Image {
	ctx := gg.NewContext(854, 456)

	bedroomImage := utils.GetAsset("8ball/backgrounds/" + data.BackgroundTheme + ".png")
	responseBoxImage := utils.GetAsset("8ball/response_boxes/" + data.BackgroundTheme + ".png")
	textBoxImage := utils.GetAsset("8ball/text_boxes/" + data.TextBoxTheme + ".png")
	menheraImage := utils.GetAsset(getRandomMenheraImage(data.Type, data.MenheraTheme))

	ctx.DrawImage(bedroomImage, 0, 0)
	ctx.DrawImage(menheraImage, 10, 10)
	ctx.DrawImage(textBoxImage, 40, 250)
	ctx.DrawImage(responseBoxImage, 440, 20)

	fontSize := 22
	if len(data.Username) <= 20 {
		fontSize = 28
	}

	ctx.SetFontFace(*utils.GetFont("Sans", float64(fontSize)))

	ctx.SetHexColor(getUsernameTextColor(data.TextBoxTheme))
	ctx.DrawStringAnchored(data.Username, 440, 339, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("Sans", 36))

	question := data.Question

	if !strings.HasSuffix(data.Question, "?") {
		question = data.Question + "?"
	}

	ctx.SetHexColor(getQuestionTextColor(data.TextBoxTheme))
	ctx.DrawStringWrapped(question, 440, 380, 0.5, 0.5, 700, 1, 1)

	ctx.SetHexColor(getResponseTextColor(data.BackgroundTheme))
	ctx.DrawStringWrapped(data.Answer, 645, 140, 0.5, 0.5, 360, 1, 1)

	return ctx.Image()
}
