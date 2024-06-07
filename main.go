package main

import (
	"fmt"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"github.com/ImArnav19/go-fiber-crm/database"
	"github.com/ImArnav19/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLead) //functions to be created
	app.Get("/api/v1/lead/:id", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated!")

}

func main() {
	app := fiber.New() //fiber instance
	initDB()
	setupRoutes(app)
	app.Listen(8000)
	// defer database.DBConn.Close() //defer used just to delay execution , let it execute at last

}
