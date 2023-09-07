package pet

import (
	"fakestore_go/entity"
	"fakestore_go/repository/pet"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Controller struct {
	logger *logrus.Logger
	repo   *pet.Repository
}

func New(logger *logrus.Logger, repo *pet.Repository) *Controller {
	return &Controller{
		logger: logger,
		repo:   repo,
	}
}

func (c *Controller) Register(app *fiber.App) {
	g := app.Group("/pet")
	g.Post("", c.InsertNewPet)
	g.Get("/spawn", c.SpawnLog)
	g.Get("/:id", c.GetPet)
}

func (c *Controller) InsertNewPet(ctx *fiber.Ctx) error {
	req := entity.Pet{}
	err := ctx.BodyParser(&req)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	resp, err := c.repo.InsertNewPetToDatabase(req)
	if err != nil {
		c.logger.Error(err)
		return ctx.Status(http.StatusBadRequest).JSON(resp)
	}
	return ctx.JSON(resp)
}

func (c *Controller) GetPet(ctx *fiber.Ctx) error {
	strPetID := ctx.Params("id")
	petID, err := strconv.ParseInt(strPetID, 10, 64)
	if err != nil {
		c.logger.Error(err)
		return ctx.Status(http.StatusBadRequest).SendString(fmt.Sprintf("%s is not an integer", strPetID))
	}
	resp, err := c.repo.GetPetFromDatabase(petID)
	if err != nil {
		c.logger.Error(err)
		rp, err := c.repo.GetPetByID(petID)
		if err != nil {
			c.logger.Error(err)
			return ctx.Status(http.StatusBadRequest).SendString("cannot find pet")
		}
		return ctx.JSON(rp)
	}
	return ctx.JSON(resp)
}

func (c *Controller) SpawnLog(ctx *fiber.Ctx) error {
	c.logger.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	c.logger.Warn("hello warning")
	c.logger.Error("error brow")
	return ctx.Status(http.StatusBadRequest).SendString("log spawned")
}
