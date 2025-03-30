package routes

import (
	"DeathRoll/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	informerHandler := handler.NewInformerHandler()
	informationHub := app.Group("/information")
	informationHub.Get("/", informerHandler.AddNewInformation)

	deathRoutes := informationHub.Group("/death")
	deathRoutes.Get("/", informerHandler.AddNewInformation)
	deathRoutes.Post("/", informerHandler.AddNewInformation)

	damageRoutes := informationHub.Group("/damage")
	damageRoutes.Get("/", informerHandler.AddNewInformation)
	damageRoutes.Post("/", informerHandler.AddNewInformation)

	donation := informationHub.Group("/donation")
	donation.Get("/", informerHandler.AddNewInformation)

	donator := app.Group("/donator")
	donator.Get("/", informerHandler.AddNewInformation)

	fundRaiser := app.Group("/fundraiser")
	fundRaiser.Get("/", informerHandler.AddNewInformation)
	fundRaiser.Post("/", informerHandler.AddNewInformation)
}
