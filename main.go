package main

import (
	"DeathRoll/cmd"
	"DeathRoll/internal/config"
	"DeathRoll/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	cfg := config.NewConfiguration()

	utils.Init()
	mainLogger := utils.NewLogger("main")
	mainLogger.Info("Starting app.....")
	cmd.Start(app, cfg)

}
