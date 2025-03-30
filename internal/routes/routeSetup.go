package routes

import (
	"DeathRoll/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	informerHandler := handler.NewInformerHandler()

	// Disaster Reports API
	reports := app.Group("/reports")
	{
		// Death reports
		death := reports.Group("/death")
		death.Post("/", informerHandler.ReportDeath)    // Submit death toll
		death.Get("/", informerHandler.GetDeathReports) // Get death reports

		// Damage reports
		damage := reports.Group("/damage")
		damage.Post("/", informerHandler.ReportDamage) // Submit damage assessment
		damage.Get("/", informerHandler.GetDamageReports)

		// Injury reports
		injury := reports.Group("/injury")
		injury.Post("/", informerHandler.ReportInjury) // Submit injury count
		injury.Get("/", informerHandler.GetInjuryReports)
	}

	// Donation Management
	donations := app.Group("/donations")
	{
		donations.Get("/", informerHandler.GetDonations)    // List all donations
		donations.Post("/", informerHandler.CreateDonation) // Create new donation
		donations.Get("/:id", informerHandler.GetDonationByID)
	}

	// Fundraiser Management
	fundraisers := app.Group("/fundraisers")
	{
		fundraisers.Get("/", informerHandler.GetApprovedFundraisers)   // List approved fundraisers
		fundraisers.Post("/", informerHandler.SubmitFundraiserRequest) // Submit new fundraiser request
		fundraisers.Patch("/:id/approve", informerHandler.ApproveFundraiser)
	}

	// Live Updates WebSocket
	app.Get("/ws/disaster-updates", informerHandler.HandleLiveUpdates)
}
