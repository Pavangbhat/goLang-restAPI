package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pavangbhat/gotest/controllers"
)


func Setup(app *fiber.App){

	app.Get("/goals",controllers.GetGoals)
	app.Post("/goal/add",controllers.CreateGoal)
	app.Post("/goal/update/:id",controllers.UpdateGoal)
	app.Post("/goal/delete/:id",controllers.DeleteGoal)
	
}

