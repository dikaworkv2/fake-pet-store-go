package main

import (
	"fakestore_go/controller/logger"
	"fakestore_go/controller/pet"
	petrepo "fakestore_go/repository/pet"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	instana "github.com/instana/go-sensor"
	"github.com/jmoiron/sqlx"
	logrus "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	instana.InitCollector(instana.DefaultOptions())
	lg := logrus.New()
	lg.SetLevel(logrus.WarnLevel)
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
	if _, ok := os.LookupEnv("INSTANA_DEBUG"); ok {
		lg.Level = logrus.DebugLevel
	}

	// use logrus to log the Instana Go Collector messages
	instana.SetLogger(lg)
	app := fiber.New(fiber.Config{AppName: "Fake Pet Store"})
	httpCli := http.Client{}

	// setup db
	db, err := sqlx.Open("mysql", "root:root_password@tcp(localhost:3306)/mydatabase")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		lg.Error(err)
		return
	}
	defer db.Close()
	fmt.Println("JALAN WOIIIIIIIII \n\n\n\n")
	// pet repo
	pRepo := petrepo.New(&httpCli, db)
	// pet controller
	petC := pet.New(lg, pRepo)
	petC.Register(app)

	logC := logger.New(lg)
	logC.Register(app)
	app.Listen(":3030")
}
