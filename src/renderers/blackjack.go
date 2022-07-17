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
	ctx := gg.NewContext(630, 460)

	tableImage, _ := utils.GetResizedAsset("tables/"+data.TableTheme+".png", 630, 460)

	ctx.DrawImage(tableImage, 0, 0)

	baseHexColor := GetFontColorByTableTheme(data.TableTheme)

	ctx.SetHexColor(baseHexColor)
	ctx.SetFontFace(*utils.GetFont("Impact", 36))

	ctx.DrawStringAnchored(data.I18n.DealerHand, 278, 36, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(data.MenheraTotal), 278, 78, 0.5, 0)

	ctx.DrawStringAnchored(data.I18n.YourHand, 280, 300, 0.5, 0)
	ctx.DrawStringAnchored(strconv.Itoa(data.UserTotal), 280, 336, 0.5, 0)

	ctx.DrawStringAnchored(strconv.Itoa(data.Aposta*2), 240, 240, 0, 0)

	ctx.SetHexColor(utils.ShadeColor(baseHexColor, -10))

	menheraStartW := (295 - 40*len(data.MenheraCards))
	userStartW := (295 - 40*len(data.UserCards))

	ctx.DrawRoundedRectangle(float64(menheraStartW-5), 85, float64(len(data.MenheraCards)*80+3), 97, 5)
	ctx.DrawRoundedRectangle(float64(userStartW-5), 353, float64(len(data.UserCards)*80+3), 97, 5)
	ctx.Fill()

	for i, card := range data.MenheraCards {
		var cardImage image.Image
		if card.Hidden {
			cardImage, _ = utils.GetResizedAsset("card_backgrounds/"+data.BackgroundCardTheme+".png", 72, 84)
		} else {
			cardImage, _ = utils.GetResizedAsset("cards/"+data.CardTheme+"/"+strconv.Itoa(card.Id)+".png", 72, 84)
		}
		ctx.DrawImage(cardImage, menheraStartW+(80*i), 93)
	}

	for i, card := range data.UserCards {
		cardImage, _ := utils.GetResizedAsset("cards/"+data.CardTheme+"/"+strconv.Itoa(card.Id)+".png", 72, 84)
		ctx.DrawImage(cardImage, userStartW+(80*i), 360)
	}

	return ctx.Image()
}
