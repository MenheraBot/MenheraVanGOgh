package renderers

import (
	"image"
	"strconv"

	"github.com/fogleman/gg"

	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

type BlackjackI18n struct {
	YourHand   string `json:"yourHand"`
	DealerHand string `json:"dealerHand"`
}

type BlackjackCard struct {
	Value  int  `json:"value"`
	IsAce  bool `json:"isAce"`
	Id     int  `json:"id"`
	Hidden bool `json:"hidden"`
}

type BlackjackData struct {
	UserCards           []BlackjackCard `json:"userCards"`
	MenheraCards        []BlackjackCard `json:"menheraCards"`
	UserTotal           int             `json:"userTotal"`
	MenheraTotal        int             `json:"menheraTotal"`
	I18n                BlackjackI18n   `json:"i18n"`
	Aposta              int             `json:"aposta"`
	CardTheme           string          `json:"cardTheme"`
	TableTheme          string          `json:"tableTheme"`
	BackgroundCardTheme string          `json:"backgroundCardTheme"`
}

func GetFontColorByTableTheme(theme string) string {
	switch theme {
	case "green":
		return "#2aa421"
	case "blue":
		return "#8194f9"
	case "red":
		return "#ff383a"
	case "pink":
		return "#f231b4"
	case "rounded":
		return "#79d861"
	case "gauderios":
		return "#2aa321"
	case "atemporal":
		return "#a760e6"
	default:
		return "#FFF"
	}
}

func RenderBlackjack(data *BlackjackData) image.Image {
	ctx := gg.NewContext(1080, 720)

	tableImage := utils.GetAsset("tables/" + data.TableTheme + ".png")

	ctx.DrawImage(tableImage, 0, 0)

	baseHexColor := GetFontColorByTableTheme(data.TableTheme)

	ctx.SetHexColor(baseHexColor)
	ctx.SetFontFace(*utils.GetFont("Impact", 72))

	ctx.DrawStringAnchored(data.I18n.DealerHand, 500, 65, 0.5, 0)

	ctx.SetFontFace(*utils.GetFont("Impact", 58))
	ctx.DrawStringAnchored(strconv.Itoa(data.MenheraTotal), 475, 128, 0.5, 0)

	ctx.DrawStringAnchored(data.I18n.YourHand, 490, 470, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(data.UserTotal), 480, 535, 0.5, 0)

	ctx.DrawStringAnchored(strconv.Itoa(data.Aposta*2), 420, 390, 0, 0)

	ctx.SetHexColor(utils.ShadeColor(baseHexColor, -10))

	menheraStartW := (490 - 60*len(data.MenheraCards))
	userStartW := (490 - 60*len(data.UserCards))

	ctx.DrawRoundedRectangle(float64(menheraStartW-10), 150, float64(len(data.MenheraCards)*125+15), 150, 5)
	ctx.DrawRoundedRectangle(float64(userStartW-10), 545, float64(len(data.UserCards)*125+15), 150, 5)
	ctx.Fill()

	for i, card := range data.MenheraCards {
		var cardImage image.Image
		if card.Hidden {
			cardImage = utils.GetAsset("card_backgrounds/" + data.BackgroundCardTheme + ".png")
		} else {
			cardImage = utils.GetAsset("cards/" + data.CardTheme + "/" + strconv.Itoa(card.Id) + ".png")
		}
		ctx.DrawImage(cardImage, menheraStartW+(125*i), 160)
	}

	for i, card := range data.UserCards {
		cardImage := utils.GetAsset("cards/" + data.CardTheme + "/" + strconv.Itoa(card.Id) + ".png")
		ctx.DrawImage(cardImage, userStartW+(125*i), 555)
	}

	return ctx.Image()
}
