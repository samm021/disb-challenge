package main

import (
	"disbursement-service/domain"
	"disbursement-service/internal/api"
	"disbursement-service/internal/config"
	"disbursement-service/internal/database"
	"disbursement-service/internal/middleware"
	"disbursement-service/internal/repository"
	"disbursement-service/internal/service"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.Get()
	app := fiber.New()

	// connect to DB & run initial migration
	dbConnection := database.GetDatabaseConnection(config)
	database.InitMigration(dbConnection, &domain.User{}, &domain.Transaction{})

	// instantiate repositories
	userRepository := repository.NewUser(dbConnection)
	transactionRepository := repository.NewTransaction(dbConnection)

	// instantiate services
	xenditService := service.NewXendit(config)
	userService := service.NewUser(userRepository)
	transactionService := service.NewTransaction(transactionRepository, userService, xenditService)

	// instantiate middlewares
	authMiddleware := middleware.Authenticate(userService)

	// instantiate routes
	api.NewTransaction(app, authMiddleware, transactionService)

	log.Print(app.Listen(fmt.Sprintf(":%s", config.Server.Port)))
}
