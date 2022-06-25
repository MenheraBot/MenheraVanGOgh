package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"log"

	"github.com/MenheraBot/MenheraVanGOgh/src/renderers"
	"github.com/MenheraBot/MenheraVanGOgh/src/renderers/profiles"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
)

func HandleWebsocketRequest(toRender string, msg []byte, util *utils.Utils) *string {
	var res image.Image

	switch toRender {
	case "astolfo":
		data := renderers.AstolfoData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderAstolfo(&data, util)
	case "blackjack":
		data := renderers.BlackjackData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderBlackjack(&data, util)
	case "8ball":
		data := renderers.EightballData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderEightball(&data, util)
	case "gado":
		data := renderers.GadoData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderGado(&data, util)
	case "macetava":
		data := renderers.MacetavaData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderMacetava(&data, util)
	case "philo":
		data := renderers.PhiloData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderPhilo(&data, util)
	case "ship":
		data := renderers.ShipData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderShip(&data, util)
	case "trisal":
		data := renderers.TrisalData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderTrisal(&data, util)
	case "vasco":
		data := renderers.VascoData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderVasco(&data, util)
	case "preview":
		data := renderers.PreviewData{}
		json.Unmarshal(msg, &data)
		res = renderers.RenderPreview(&data, util)
	case "profile":
		data := utils.ProfileData{}
		json.Unmarshal(msg, &data)
		switch data.Type {
		case "fortification":
			res = profiles.RenderFortification(&data.User, &data.I18n, util)
		case "warrior":
			res = profiles.RenderWarrior(&data.User, &data.I18n, util)
		case "christmas_2021":
			res = profiles.RenderChristmas(&data.User, &data.I18n, util)
		case "kawaii":
			res = profiles.RenderKawaii(&data.User, &data.I18n, util)
		case "id03":
			res = profiles.RenderID03(&data.User, &data.I18n, util)
		case "without_soul":
			res = profiles.RenderWithoutSoul(&data.User, &data.I18n, util)
		case "upsidedown":
			res = profiles.RenderUpsideDown(&data.User, &data.I18n, util)
		case "default":
			res = profiles.RenderDefault(&data.User, &data.I18n, util)
		default:
			res = profiles.RenderDefault(&data.User, &data.I18n, util)
		}
	}

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	if err != nil {
		log.Print(err)
	}

	return &encodedString
}
