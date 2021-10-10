package main

import (
	"fmt"
	"log"
	"os"

	"github.com/farukak/todos/database"
	"github.com/farukak/todos/handlers"
	"github.com/farukak/todos/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "Tasko",
		AppName:       "Tasko v1.0.0",
		ErrorHandler:  handlers.ErrorHandler,
	})

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Turkey/Istanbul",
	}))

	app.Use(cors.New())

	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many.",
			})
		},
	}))

	database.ConnectDB()

	route.Route(app)

	Addr := fmt.Sprintf("127.0.0.1:%s", os.Getenv("APP_PORT"))

	log.Fatal(app.Listen(Addr))

	//defer database.DB.Close()

}
