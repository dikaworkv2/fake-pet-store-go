package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	logger *logrus.Logger
}

func New(logger *logrus.Logger) *Controller {
	return &Controller{logger: logger}
}

func (c *Controller) Register(app *fiber.App) {
	g := app.Group("/log")
	g.Get("/info", c.SendInfoLog)
	g.Get("/warn", c.SendWarnLog)
	g.Get("/error", c.SendErrorLog)
}

func (c *Controller) SendInfoLog(ctx *fiber.Ctx) error {
	c.logger.Info("send info")
	return ctx.SendString("sent")
}

func (c *Controller) SendWarnLog(ctx *fiber.Ctx) error {
	c.logger.Warn("send warn")
	return ctx.SendString("sent")
}

func (c *Controller) SendErrorLog(ctx *fiber.Ctx) error {
	c.logger.Error("send error")
	return ctx.SendString("sent")
}
