package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pavangbhat/gotest/routes"
)
  

  func main(){
	app := fiber.New()
	 
	// Middlewares
	app.Use(logger.New())

	//routes setup
	routes.Setup(app)

	app.Listen(":3000")
  }