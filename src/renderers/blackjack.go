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

func getFontColorByTheme(theme string) string {
	switch theme {
	case "green":
		return "#2aa421"
	case "blue":
		return "#3e86e9"
	case "red":
		return "#ff383a"
	case "pink":
		return "#f231b4"
	case "rounded":
		return "#2aa421"
	case "gauderios":
		return "#2aa321"
	case "atemporal":
		return "#a760e6"
	default:
		return "#FFF"
	}
}

func RenderBlackjack(data *BlackjackData, util utils.Utils) image.Image {
	ctx := gg.NewContext(630, 460)

	tableImage, _ := util.GetResizedAsset("tables/"+data.TableTheme+".png", 630, 460)

	ctx.DrawImage(tableImage, 0, 0)

	baseHexColor := getFontColorByTheme(data.TableTheme)

	ctx.SetHexColor(baseHexColor)
	ctx.LoadFontFace(util.GetFontPath("Impact"), 36)

	dealerText := data.I18n.DealerHand + "\n" + strconv.Itoa(data.MenheraTotal)
	userText := data.I18n.YourHand + "\n" + strconv.Itoa(data.UserTotal)

	util.FillStrokedText(ctx, dealerText, 278, 36, 1000, 1000, 35, 2, "#000", baseHexColor, 0.5)
	util.FillStrokedText(ctx, userText, 280, 300, 1000, 1000, 40, 2, "#000", baseHexColor, 0.5)

	util.StrokeText(ctx, strconv.Itoa(data.Aposta*2), 240, 240, 2, "#000", "#FFFF00", 0)

	ctx.SetHexColor(util.ShadeColor(baseHexColor, -10))

	menheraStartW := (295 - 40*len(data.MenheraCards))
	userStartW := (295 - 40*len(data.UserCards))

	ctx.DrawRoundedRectangle(float64(menheraStartW-5), 85, float64(len(data.MenheraCards)*80+3), 97, 5)
	ctx.DrawRoundedRectangle(float64(userStartW-5), 353, float64(len(data.UserCards)*80+3), 97, 5)
	ctx.Fill()

	for i, card := range data.MenheraCards {
		var cardImage image.Image
		if card.Hidden {
			cardImage, _ = util.GetResizedAsset("card_backgrounds/" + data.BackgroundCardTheme + ".png", 72, 84)
		} else {
			cardImage, _ = util.GetResizedAsset("cards/"+data.CardTheme+"/"+strconv.Itoa(card.Id)+".png", 72, 84)
		}
		ctx.DrawImage(cardImage, menheraStartW+(80*i), 93)
	}

	for i, card := range data.UserCards {
		cardImage, _ := util.GetResizedAsset("cards/"+data.CardTheme+"/"+strconv.Itoa(card.Id)+".png", 72, 84)
		ctx.DrawImage(cardImage, userStartW+(80*i), 360)
	}

	return ctx.Image()
}
