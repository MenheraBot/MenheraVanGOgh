package controllers

import (
	"image/png"

	"github.com/MenheraBot/MenheraVanGOgh/src/renderers"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/gofiber/fiber/v2"
)

var encoder = png.Encoder{
	CompressionLevel: png.BestSpeed,
}

var utilities = utils.New()

func Astolfo(c *fiber.Ctx) error {
	data := new(renderers.AstolfoData)

	c.BodyParser(data)

	res := renderers.RenderAstolfo(data, utilities)
	
	return encoder.Encode(c.Context(), res)
}

func Philo(c *fiber.Ctx) error {
	data := new(renderers.PhiloData)

	c.BodyParser(data)

	res := renderers.RenderPhilo(data, utilities)
	
	return encoder.Encode(c.Context(), res)
}

func Trisal(c *fiber.Ctx) error {
	data := new(renderers.TrisalData)

	c.BodyParser(data)

	res := renderers.RenderTrisal(data, utilities)
	
	return encoder.Encode(c.Context(), res)
}

func Ship(c *fiber.Ctx) error {
	data := new(renderers.ShipData)

	c.BodyParser(data)

	res := renderers.RenderShip(data, utilities)
	
	return encoder.Encode(c.Context(), res)
}
