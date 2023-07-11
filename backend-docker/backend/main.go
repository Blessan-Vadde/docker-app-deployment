package main

import (
	"employees/controller"
	"employees/repository"
	"employees/routes"
	"employees/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "true",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	db := initializeDatabaseConnection()
	repository.RunMigrations(db)
	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)
	routes.RegisterRoute(app, employeeController)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalln(fmt.Sprintf("error starting the server %s", err.Error()))
	}
}
func initializeDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  createDsn(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(fmt.Sprintf("error connecting with database %s", err.Error()))
	}

	return db
}

func createDsn() string {
	dsnFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dbHost := "dev-database.cx2v0znhuvnq.eu-central-1.rds.amazonaws.com"
	dbUser := "postgres"
	dbPassword := "adminpassword"
	dbName := "nspglobal"
	dbPort := "5432"
	return fmt.Sprintf(dsnFormat, dbHost, dbUser, dbPassword, dbName, dbPort)
}
