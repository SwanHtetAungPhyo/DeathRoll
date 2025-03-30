package cmd

import (
	"DeathRoll/internal/config"
	"DeathRoll/internal/routes"
	"DeathRoll/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

func Start(app *fiber.App, cfg *config.Configuration) {
	cmdLogger := utils.NewLogger("Start")
	routes.SetupRoutes(app)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			cmdLogger.Fatal(err.Error())
		}
	}()

	osChan := make(chan os.Signal, 3)
	signal.Notify(osChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-osChan
	cmdLogger.Info("Shutting down...")
	if err := app.Shutdown(); err != nil {
		cmdLogger.Fatal(err.Error())
	}
}
