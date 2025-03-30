package handler

import "github.com/gofiber/fiber/v2"

func (i *InformerHandler) ReportDeath(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (i *InformerHandler) GetDeathReports(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"Current Death": 1200,
	})
}
