package main

import (
	"DeathRoll/cmd"
	"DeathRoll/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	cfg := config.NewConfiguration()

	cmd.Start(app, cfg)

}
