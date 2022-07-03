package controllers

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"

	"github.com/MenheraBot/MenheraVanGOgh/src/renderers"
	"github.com/MenheraBot/MenheraVanGOgh/src/renderers/profiles"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/gin-gonic/gin"
)

var encoder = png.Encoder{
	CompressionLevel: png.BestSpeed,
}

func Astolfo(c *gin.Context) {
	data := new(renderers.AstolfoData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderAstolfo(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Philo(c *gin.Context) {
	data := new(renderers.PhiloData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderPhilo(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Trisal(c *gin.Context) {
	data := new(renderers.TrisalData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderTrisal(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Ship(c *gin.Context) {
	data := new(renderers.ShipData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderShip(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Gado(c *gin.Context) {
	data := new(renderers.GadoData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderGado(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Macetava(c *gin.Context) {
	data := new(renderers.MacetavaData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderMacetava(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Eightball(c *gin.Context) {
	data := new(renderers.EightballData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderEightball(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Vasco(c *gin.Context) {
	data := new(renderers.VascoData)

	err := c.BindJSON(data)
	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderVasco(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Blackjack(c *gin.Context) {
	data := new(renderers.BlackjackData)

	err := c.BindJSON(data)
	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderBlackjack(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Profile(c *gin.Context) {
	data := new(utils.ProfileData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	var res image.Image

	switch data.Type {
	case "fortification":
		res = profiles.RenderFortification(&data.User, &data.I18n)
	case "warrior":
		res = profiles.RenderWarrior(&data.User, &data.I18n)
	case "christmas_2021":
		res = profiles.RenderChristmas(&data.User, &data.I18n)
	case "kawaii":
		res = profiles.RenderKawaii(&data.User, &data.I18n)
	case "id03":
		res = profiles.RenderID03(&data.User, &data.I18n)
	case "without_soul":
		res = profiles.RenderWithoutSoul(&data.User, &data.I18n)
	case "upsidedown":
		res = profiles.RenderUpsideDown(&data.User, &data.I18n)
	case "default":
		res = profiles.RenderDefault(&data.User, &data.I18n)
	default:
		res = profiles.RenderDefault(&data.User, &data.I18n)
	}

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Preview(c *gin.Context) {
	data := new(renderers.PreviewData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderPreview(data)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}
