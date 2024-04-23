package main

import (
	"fmt"
	"log"
	user_handler "main/internal/handler/rest/user"
	user_repo "main/internal/repo/user"
	user_uc "main/internal/usecase/user"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// SETUP DATABASE
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=allow TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// REPO INITIALIZATION
	userRepo := user_repo.New(db)

	// USECASE INITIALIZATION
	userUsecase := user_uc.New(userRepo)

	// HANDLER INITIALIZATION
	userHandler := user_handler.New(userUsecase)

	// SETUP ROUTER
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/DisplayAllUsers", userHandler.DisplayAllUsers)
	app.Post("/DisplayUser", userHandler.DisplayUser)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Printf("server listening at %s", serverPort)

	app.Listen(serverPort)
}
