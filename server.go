package main

import (
	"log"

	"github.com/dickanirwansyah/blogspot/database"
	"github.com/dickanirwansyah/blogspot/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file")
	}

	database.ConnectDB()

}

func main() {

	database.ConnectDB()

	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

	defer sqlDb.Close()

	app := fiber.New()

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")
}
