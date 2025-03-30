package handler

import (
	"DeathRoll/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	InformerHandlerMethod interface {
		AddNewInformation(c *fiber.Ctx) error
	}
	InformerHandler struct {
		log *logrus.Logger
	}
)

func NewInformerHandler() *InformerHandler {
	return &InformerHandler{log: utils.NewLogger("InformerHandler")}
}

func (i *InformerHandler) AddNewInformation(c *fiber.Ctx) error {
	panic("implement me")
}

func (i *InformerHandler) CreateDonation(ctx *fiber.Ctx) error {
	panic("implement me")
}
