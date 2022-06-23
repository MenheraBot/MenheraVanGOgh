package controllers

import (
	"image"
	"image/png"

	"github.com/MenheraBot/MenheraVanGOgh/src/renderers"
	"github.com/MenheraBot/MenheraVanGOgh/src/renderers/profiles"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/gin-gonic/gin"
)

var encoder = png.Encoder{
	CompressionLevel: png.BestSpeed,
}

var utilities = utils.New()

func Astolfo(c *gin.Context) {
	data := new(renderers.AstolfoData)

	c.BindJSON(data)

	res := renderers.RenderAstolfo(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Philo(c *gin.Context) {
	data := new(renderers.PhiloData)

	c.BindJSON(data)

	res := renderers.RenderPhilo(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Trisal(c *gin.Context) {
	data := new(renderers.TrisalData)

	c.BindJSON(data)

	res := renderers.RenderTrisal(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Ship(c *gin.Context) {
	data := new(renderers.ShipData)

	c.BindJSON(data)

	res := renderers.RenderShip(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Gado(c *gin.Context) {
	data := new(renderers.GadoData)

	c.BindJSON(data)

	res := renderers.RenderGado(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Macetava(c *gin.Context) {
	data := new(renderers.MacetavaData)

	c.BindJSON(data)

	res := renderers.RenderMacetava(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Eightball(c *gin.Context) {
	data := new(renderers.EightballData)

	c.BindJSON(data)

	res := renderers.RenderEightball(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Vasco(c *gin.Context) {
	data := new(renderers.VascoData)

	c.BindJSON(data)

	res := renderers.RenderVasco(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Blackjack(c *gin.Context) {
	data := new(renderers.BlackjackData)

	c.BindJSON(data)

	res := renderers.RenderBlackjack(data, utilities)

	encoder.Encode(c.Writer, res)
}

func Profile(c *gin.Context) {
	data := new(utils.ProfileData)

	c.BindJSON(data)

	var res image.Image

	switch data.Type {

	case "fortification":
		res = profiles.RenderFortification(&data.User, &data.I18n, utilities)
	case "warrior":
		res = profiles.RenderWarrior(&data.User, &data.I18n, utilities)
	case "christmas_2021":
		res = profiles.RenderChristmas(&data.User, &data.I18n, utilities)
	case "kawaii":
		res = profiles.RenderKawaii(&data.User, &data.I18n, utilities)
	case "id03":
		res = profiles.RenderID03(&data.User, &data.I18n, utilities)
	case "without_soul":
		res = profiles.RenderWithoutSoul(&data.User, &data.I18n, utilities)
	case "upsidedown":
		res = profiles.RenderUpsideDown(&data.User, &data.I18n, utilities)
	case "default":
		res = profiles.RenderDefault(&data.User, &data.I18n, utilities)
	default:
		res = profiles.RenderDefault(&data.User, &data.I18n, utilities)
	}

	encoder.Encode(c.Writer, res)
}
