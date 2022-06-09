package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hashcott/go-project/go-fiber/config"
	"github.com/hashcott/go-project/go-fiber/entity"
	"log"
)

func initializeDB() {
	if err := config.InitDB(); err != nil {
		panic("Can't connect to config")
	} else {
		fmt.Println("Connected to DB.")
		config.DB.AutoMigrate(entity.Lead{})
	}
}

func main() {
	app := fiber.New()
	initializeDB()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world !")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error occurred:", err)
	}
}
