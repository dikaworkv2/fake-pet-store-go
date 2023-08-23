package main

import (
	"fakestore_go/controller/pet"
	petrepo "fakestore_go/repository/pet"
	"github.com/gofiber/fiber/v2"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	lg := logrus.New()
	lg.SetLevel(logrus.DebugLevel)
	lg.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		// Customize colors for different log levels
		DisableColors:    false,
		DisableTimestamp: false,
		QuoteEmptyFields: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
	})
	app := fiber.New(fiber.Config{AppName: "Fake Pet Store"})
	httpCli := http.Client{}

	// pet repo
	pRepo := petrepo.New(&httpCli)
	// pet controller
	petC := pet.New(lg, pRepo)
	petC.Register(app)

	app.Listen(":3030")
}
