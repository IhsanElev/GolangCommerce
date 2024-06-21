package main

import (
	"GolangEcommerceDDD/app/auth"
	"GolangEcommerceDDD/app/product"
	"GolangEcommerceDDD/app/transaction"
	"GolangEcommerceDDD/external/database"
	infrafiber "GolangEcommerceDDD/infra/fiber"
	"GolangEcommerceDDD/internal/config"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	if db != nil {
		fmt.Println("db connected")
	}

	router := fiber.New(
		fiber.Config{
			Prefork: true,
			AppName: config.Cfg.App.Name,
		})
	router.Use(infrafiber.Trace())
	auth.Init(router, db)
	product.Init(router, db)
	transaction.Init(router, db)
	router.Listen(config.Cfg.App.Port)

}
