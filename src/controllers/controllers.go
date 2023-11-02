package controllers

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"

	"github.com/MenheraBot/MenheraVanGOgh/src/database"
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

func Trisal(c *gin.Context, db *database.Database) {
	data := new(renderers.TrisalData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderTrisal(data, db)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Ship(c *gin.Context, db *database.Database) {
	data := new(renderers.ShipData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderShip(data, db)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Gado(c *gin.Context, db *database.Database) {
	data := new(renderers.GadoData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderGado(data, db)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}

func Macetava(c *gin.Context, db *database.Database) {
	data := new(renderers.MacetavaData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderMacetava(data, db)

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

func Vasco(c *gin.Context, db *database.Database) {
	data := new(renderers.VascoData)

	err := c.BindJSON(data)
	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderVasco(data, db)

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

func Profile(c *gin.Context, db *database.Database) {
	data := new(utils.ProfileData)

	err := c.BindJSON(data)

	if err != nil {
		log.Print(err)
	}

	base64Profile, err := db.GetCachedProfileImage(data.User.Id, data.HashedData)

	if err == nil {
		c.String(200, base64Profile)
		return
	}

	var res image.Image

	switch data.Type {
	case "fortification":
		res = profiles.RenderFortification(&data.User, &data.I18n, data.CustomEdits, db)
	case "warrior":
		res = profiles.RenderWarrior(&data.User, &data.I18n, data.CustomEdits, db)
	case "christmas_2021":
		res = profiles.RenderChristmas(&data.User, &data.I18n, data.CustomEdits, db)
	case "kawaii":
		res = profiles.RenderKawaii(&data.User, &data.I18n, db)
	case "id03":
		res = profiles.RenderID03(&data.User, &data.I18n, db)
	case "without_soul":
		res = profiles.RenderWithoutSoul(&data.User, &data.I18n, db)
	case "upsidedown":
		res = profiles.RenderUpsideDown(&data.User, &data.I18n, db)
	case "default":
		res = profiles.RenderDefault(&data.User, &data.I18n, db)
	case "gatito":
		res = profiles.RenderGatito(&data.User, &data.I18n, db)
	case "mural":
		res = profiles.RenderPersonalSpace(&data.User, &data.I18n, data.CustomEdits, db)
	case "hello_kitty":
		res = profiles.RenderHelloKitty(&data.User, &data.I18n, db)
	case "sunflower":
		res = profiles.RenderSunflower(&data.User, &data.I18n, db)
	case "gallery":
		res = profiles.RenderGallery(&data.User, &data.I18n, db)
	default:
		res = profiles.RenderDefault(&data.User, &data.I18n, db)
	}

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	stringedImage := base64.StdEncoding.EncodeToString(buff.Bytes())

	db.SetCachedProfileImage(data.User.Id, data.HashedData, stringedImage)

	c.String(200, stringedImage)
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

func Poker(c *gin.Context, db *database.Database) {
	query := c.Request.URL.Query().Get("player")

	if query == "true" {
		data := new(renderers.PokerHandData)

		err := c.BindJSON(data)
		if err != nil {
			log.Print(err)
		}

		res := renderers.RenderPokerHand(data, db)

		buff := new(bytes.Buffer)
		err = encoder.Encode(buff, res)

		if err != nil {
			log.Print(err)
		}

		c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
		return
	}

	data := new(renderers.PokerTableData)

	err := c.BindJSON(data)
	if err != nil {
		log.Print(err)
	}

	res := renderers.RenderPokerTable(data, db)

	buff := new(bytes.Buffer)
	err = encoder.Encode(buff, res)

	if err != nil {
		log.Print(err)
	}

	c.String(200, base64.StdEncoding.EncodeToString(buff.Bytes()))
}
