package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sparsh459/REST-API-with-GO-FIber/database"
	"github.com/sparsh459/REST-API-with-GO-FIber/routes"
)

// *fiber.Ctx is called context and it is everything that comes with an request

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the api test run")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)

	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	// // Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)

	// // Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
}

func main() {

	// connecting to db
	database.ConnectDb()

	// creating instance of the app
	app := fiber.New()

	setupRoutes(app)

	// start app port
	// using log.Fatal so that it prints out the current status of our server
	log.Fatal(app.Listen(":3000"))
}
