package controllers

import (
	"bytes"
	"encoding/base64"
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

func Astolfo(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.AstolfoData)

	c.BindJSON(data)

	res := renderers.RenderAstolfo(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Philo(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.PhiloData)

	c.BindJSON(data)

	res := renderers.RenderPhilo(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Trisal(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.TrisalData)

	c.BindJSON(data)

	res := renderers.RenderTrisal(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Ship(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.ShipData)

	c.BindJSON(data)

	res := renderers.RenderShip(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Gado(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.GadoData)

	c.BindJSON(data)

	res := renderers.RenderGado(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Macetava(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.MacetavaData)

	c.BindJSON(data)

	res := renderers.RenderMacetava(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Eightball(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.EightballData)

	c.BindJSON(data)

	res := renderers.RenderEightball(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Vasco(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.VascoData)

	c.BindJSON(data)

	res := renderers.RenderVasco(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Blackjack(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.BlackjackData)

	c.BindJSON(data)

	res := renderers.RenderBlackjack(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Profile(c *gin.Context, Utilities *utils.Utils) {
	data := new(utils.ProfileData)

	c.BindJSON(data)

	var res image.Image

	switch data.Type {

	case "fortification":
		res = profiles.RenderFortification(&data.User, &data.I18n, Utilities)
	case "warrior":
		res = profiles.RenderWarrior(&data.User, &data.I18n, Utilities)
	case "christmas_2021":
		res = profiles.RenderChristmas(&data.User, &data.I18n, Utilities)
	case "kawaii":
		res = profiles.RenderKawaii(&data.User, &data.I18n, Utilities)
	case "id03":
		res = profiles.RenderID03(&data.User, &data.I18n, Utilities)
	case "without_soul":
		res = profiles.RenderWithoutSoul(&data.User, &data.I18n, Utilities)
	case "upsidedown":
		res = profiles.RenderUpsideDown(&data.User, &data.I18n, Utilities)
	case "default":
		res = profiles.RenderDefault(&data.User, &data.I18n, Utilities)
	default:
		res = profiles.RenderDefault(&data.User, &data.I18n, Utilities)
	}

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))

}

func Preview(c *gin.Context, Utilities *utils.Utils) {
	data := new(renderers.PreviewData)

	c.BindJSON(data)

	res := renderers.RenderPreview(data, Utilities)

	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, res)

	if err != nil {
		panic(err)
	}

	c.Writer.WriteString(base64.StdEncoding.EncodeToString(buff.Bytes()))
}
